package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

func ReceiveSucceededWork() error {
	sub, err := global.MQ.Subscribe(commonConsts.MQSETTINGS_SucceededChannelName, handleSucceeded)
	if err != nil {
		global.Logger.Error("Failed to subscribe to MQ succeeded queue with error: ", err)
		return err
	}

	defer sub.Drain() // Ignore errors

	return err
}

func handleSucceeded(m *nats.Msg) {
	var workSucceeded commonTypes.WorkSucceeded
	if err := json.Unmarshal(m.Data, &workSucceeded); err != nil {
		global.Logger.Error("Unable to parse succeeded work: ", m.Data)
	} else {
		// Parse successfully
		// Parse feeds
		var feeds []models.Feed
		for _, rawFeed := range workSucceeded.Feeds {
			feed := models.Feed{
				Feed: types.Feed{
					AccountID:   workSucceeded.AccountID,
					Platform:    workSucceeded.Platform,
					CollectedAt: workSucceeded.DispatchAt,
					RawFeed:     rawFeed,
				},
			}
			feeds = append(feeds, feed)
		}

		// Find account
		var account models.Account
		global.DB.First(&account, workSucceeded.AccountID)

		// Update account
		interv := workSucceeded.SucceededAt.Sub(account.LastUpdated)
		platform := commonConsts.SUPPORTED_PLATFORM[workSucceeded.Platform]
		if interv < platform.MinRefreshGap {
			interv = platform.MinRefreshGap
		} else if interv > platform.MaxRefreshGap {
			interv = platform.MaxRefreshGap
		}
		account.LastUpdated = workSucceeded.SucceededAt
		account.UpdateInterval = interv

		if err := global.DB.Transaction(func(tx *gorm.DB) error {
			// do some database operations in the transaction (use 'tx' from this point, not 'db')
			platformSpecifiedFeed := models.Feed{
				Feed: types.Feed{
					Platform: workSucceeded.Platform,
				},
			}

			// Insert feeds
			tx.Scopes(models.FeedTable(platformSpecifiedFeed)).Create(&feeds)

			// Update account
			tx.Save(&account)

			// return nil will commit the whole transaction
			return nil
		}); err != nil {
			global.Logger.Error("Unable to save feeds: ", workSucceeded)
		} else {
			// Succeeded
			// Clear cache
			accountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts", account.CrossbellCharacter)
			feedsCacheKey := fmt.Sprintf("%s:%s:%d", consts.CACHE_PREFIX, "feeds", account.ID)
			clearCacheCtx := context.Background()
			global.Redis.Del(clearCacheCtx, accountsCacheKey) // To flush account update time
			global.Redis.Del(clearCacheCtx, feedsCacheKey)    // To flush cached feeds

			// Update metrics
			global.MetricsSucceededWorkCount.Inc(1)
		}
	}
}
