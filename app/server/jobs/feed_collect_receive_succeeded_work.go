package jobs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
	"sort"
)

func FeedCollectStartReceiveSucceededWork() error {
	_, err := global.MQ.Subscribe(commonConsts.MQSETTINGS_SucceededChannelName, feedCollectHandleSucceeded)
	if err != nil {
		global.Logger.Error("Failed to subscribe to MQ Feeds Collect succeeded queue with error: ", err.Error())
		return err
	}

	//defer sub.Drain() // Ignore errors

	return nil
}

func feedCollectHandleSucceeded(m *nats.Msg) {
	global.Logger.Debug("New succeeded Collect work received: ", string(m.Data))

	var workSucceeded commonTypes.WorkSucceeded
	if err := json.Unmarshal(m.Data, &workSucceeded); err != nil {
		global.Logger.Error("Unable to parse succeeded work: ", string(m.Data))
	} else {
		// Parse successfully
		// Parse feeds
		var feeds models.FeedsArray
		if len(workSucceeded.Feeds) > 0 {
			for _, rawFeed := range workSucceeded.Feeds {
				feed := models.Feed{
					Feed: types.Feed{
						AccountID:   workSucceeded.AccountID,
						Platform:    workSucceeded.Platform,
						CollectedAt: workSucceeded.DispatchAt,
						RawFeed:     rawFeed,
					},
				}
				for _, media := range rawFeed.Media {
					feed.MediaIPFSUris = append(feed.MediaIPFSUris, media.IPFSUri)
				}
				feeds = append(feeds, feed)
			}
		}

		// Sort feeds by publish time (ASC)
		sort.Sort(feeds)

		// Find account
		var account models.Account
		global.DB.First(&account, workSucceeded.AccountID)

		// Update account
		interv := workSucceeded.NewInterval
		platform := commonConsts.SUPPORTED_PLATFORM[workSucceeded.Platform]
		if interv < platform.MinRefreshGap {
			interv = platform.MinRefreshGap
		} else if interv > platform.MaxRefreshGap {
			interv = platform.MaxRefreshGap
		}
		account.LastUpdated = workSucceeded.SucceededAt
		account.UpdateInterval = interv
		account.NextUpdate = account.LastUpdated.Add(account.UpdateInterval)
		account.FeedsCount += uint(len(feeds))
		// NotesCount increase need feeds to be published on chain

		// Update character
		var character models.Character
		global.DB.First(&character, "crossbell_character_id = ?", account.CrossbellCharacterID)

		if err := global.DB.Transaction(func(tx *gorm.DB) error {
			// do some database operations in the transaction (use 'tx' from this point, not 'db')
			if len(feeds) > 0 {

				// Insert feeds
				if err := tx.Scopes(models.FeedTable(models.Feed{
					Feed: types.Feed{
						Platform: workSucceeded.Platform,
					},
				})).Create(&feeds).Error; err != nil {
					return err
				}

				// Insert medias (Can only be processed here because we need feed IDs to identify them)
				mediaMap := make(map[string]models.Media)
				for _, feed := range feeds {
					for _, media := range feed.Media {
						var singleMedia models.Media
						var ok bool
						if singleMedia, ok = mediaMap[media.IPFSUri]; !ok {
							// Try to find in database
							if err = global.DB.First(&singleMedia, "ipfs_uri = ?", media.IPFSUri).Error; errors.Is(err, gorm.ErrRecordNotFound) {
								singleMedia = models.Media{
									ID:                   0,
									CrossbellCharacterID: account.CrossbellCharacterID,
									Media:                media,
								}
							}
						}
						singleMedia.RelatedFeeds = append(singleMedia.RelatedFeeds, models.MediaFeedRecord{
							Platform: workSucceeded.Platform,
							ID:       feed.ID,
						})
						mediaMap[media.IPFSUri] = singleMedia
					}
				}

				var mediaUpdateList []models.Media
				var mediaCreateList []models.Media
				for _, singleMedia := range mediaMap {
					if singleMedia.ID == 0 {
						mediaCreateList = append(mediaCreateList, singleMedia)

						account.MediaUsage += singleMedia.FileSize
						character.MediaUsage += singleMedia.FileSize
					} else {
						mediaUpdateList = append(mediaUpdateList, singleMedia)
					}
				}

				if len(mediaUpdateList) > 0 {
					if err := tx.Save(&mediaUpdateList).Error; err != nil {
						return err
					}
				}

				if len(mediaCreateList) > 0 {
					if err := tx.Create(&mediaCreateList).Error; err != nil {
						return err
					}
				}
			}

			// Update account
			if err := tx.Save(&account).Error; err != nil {
				return err
			}

			// Update character
			if err := tx.Save(&character).Error; err != nil {
				return err
			}

			// return nil will commit the whole transaction
			return nil
		}); err != nil {
			global.Logger.Error("Unable to save feeds: ", workSucceeded)
		} else {
			// Succeeded
			// Clear cache
			accountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts", account.CrossbellCharacterID)
			feedsCacheKey := fmt.Sprintf("%s:%s:%d", consts.CACHE_PREFIX, "feeds", account.ID)
			mediasCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "medias", account.CrossbellCharacterID)
			clearCacheCtx := context.Background()
			global.Redis.Del(clearCacheCtx, accountsCacheKey) // To flush account update time
			global.Redis.Del(clearCacheCtx, feedsCacheKey)    // To flush cached feeds
			global.Redis.Del(clearCacheCtx, mediasCacheKey)   // To flush cached media list

			// Post feeds On Chain
			feedOnChainDispatchWork(&account, feeds)

			// Update metrics
			global.Metrics.Work.Succeeded.Inc(1)
		}
	}
}
