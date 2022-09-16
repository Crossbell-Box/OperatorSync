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
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"sort"
	"time"
)

func FeedCollectStartRetrieveWork() error {

	// Prepare channel

	ch, err := commonGlobal.MQ.Channel()
	if err != nil {
		global.Logger.Error("Failed to open MQ Feeds Collect retrieve channel with error: ", err.Error())
		return err
	}

	// Prepare queue

	q, err := ch.QueueDeclare(
		commonConsts.MQSETTINGS_FeedCollectRetrieveQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to prepare MQ Feeds Collect retrieve queue with error: ", err.Error())
		return err
	}

	deliveries, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Error("Failed to subscribe to MQ Feeds Collect retrieve queue with error: ", err.Error())
		return err
	}

	global.Logger.Debug("Feed Collect start listen on retrieve works...")

	go func() {
		for d := range deliveries {
			switch d.Headers[commonConsts.MQSETTINGS_FeedCollectIdentifierField] {
			case commonConsts.MQSETTINGS_FeedCollectSucceededIdentifier:
				feedCollectHandleSucceeded(&d)
			case commonConsts.MQSETTINGS_FeedCollectFailedIdentifier:
				feedCollectHandleFailed(&d)
			default:
				global.Logger.Errorf("Undefined message received: %v", d)
			}
		}
	}()

	return nil
}

func feedCollectHandleSucceeded(d *amqp.Delivery) {
	global.Logger.Debug("New succeeded Collect work received: ", string(d.Body))

	var workSucceeded commonTypes.WorkSucceeded
	if err := json.Unmarshal(d.Body, &workSucceeded); err != nil {
		global.Logger.Error("Unable to parse succeeded work: ", string(d.Body))
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
		var interv time.Duration
		if len(feeds) > 0 {
			interv = workSucceeded.NewInterval
		} else {
			// workSucceeded.NewInterval == account.NextUpdate - account.LastUpdated
			interv = workSucceeded.SucceededAt.Sub(account.LastUpdated) // So we can double the interval as nothing found since then
		}
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
		//var character models.Character
		//global.DB.First(&character, "crossbell_character_id = ?", account.CrossbellCharacterID)

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
				mediaTypedUsageInc := make(map[string]uint)
				for _, singleMedia := range mediaMap {
					if singleMedia.ID == 0 {
						mediaCreateList = append(mediaCreateList, singleMedia)

						// Record ContentType with increment
						if _, ok := mediaTypedUsageInc[singleMedia.ContentType]; ok {
							mediaTypedUsageInc[singleMedia.ContentType] += singleMedia.FileSize
						} else {
							mediaTypedUsageInc[singleMedia.ContentType] = singleMedia.FileSize
						}
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

				// Save increments to Account
				//// Find already exist ContentTypes
				for index, mediaUsageWithTypeInAccount := range account.MediaUsage {
					if increment, ok := mediaTypedUsageInc[mediaUsageWithTypeInAccount.ContentType]; ok {
						// Already exists, add to record
						account.MediaUsage[index].Usage += increment
						// Delete key
						delete(mediaTypedUsageInc, mediaUsageWithTypeInAccount.ContentType)
					}
				}
				//// Create new ContentTypes
				for contentType, increment := range mediaTypedUsageInc {
					account.MediaUsage = append(account.MediaUsage, types.MediaTypeRecord{
						ContentType: contentType,
						Usage:       increment,
					})
				}
			}

			// Update account
			if err := tx.Save(&account).Error; err != nil {
				return err
			}

			//// Update character
			//if err := tx.Save(&character).Error; err != nil {
			//	return err
			//}

			// return nil will commit the whole transaction
			return nil
		}); err != nil {
			global.Logger.Error("Unable to save feeds: ", workSucceeded)
		} else {
			// Succeeded
			// Clear cache
			accountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:list", account.CrossbellCharacterID)
			feedsCacheKey := fmt.Sprintf("%s:%s:%d", consts.CACHE_PREFIX, "feeds", account.ID)
			mediasCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "medias", account.CrossbellCharacterID)
			clearCacheCtx := context.Background()
			global.Redis.Del(clearCacheCtx, accountsCacheKey) // To flush account update time
			global.Redis.Del(clearCacheCtx, feedsCacheKey)    // To flush cached feeds
			global.Redis.Del(clearCacheCtx, mediasCacheKey)   // To flush cached media list

			if len(feeds) > 0 {
				// Post feeds On Chain
				feedOnChainDispatchWork(&account, feeds)
			}

			// Update metrics
			global.Metrics.Work.Succeeded.Inc(1)
		}
	}
}

func feedCollectHandleFailed(d *amqp.Delivery) {
	global.Logger.Warn("New failed Collect work received: ", string(d.Body))

	var workFailed commonTypes.WorkFailed
	if err := json.Unmarshal(d.Body, &workFailed); err != nil {
		global.Logger.Error("Unable to parse failed work: ", string(d.Body))
	} else {
		global.Logger.Warn("Work failed for: ", workFailed)
		global.Metrics.Work.Failed.Inc(1)
	}
}