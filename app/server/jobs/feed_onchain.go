package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func feedOnChainDispatchWork(account *models.Account, feeds []models.Feed) {
	if utils.IsAccountOnChainPaused(account) {
		global.Logger.Errorf("OnChain process has been paused for account %s#%d with reason: %s", account.Platform, account.ID, account.OnChainPauseMessage)
		return
	}

	for index, feed := range feeds {

		ipfsUri, tx, err := OneFeedOnChain(account, &feed)
		if err != nil {
			global.Logger.Errorf("Failed to push all feeds on chain with error: %s", err.Error())
			break
		} else {
			// `feed` is read-only and cannot change its value
			feeds[index].IPFSUri = ipfsUri
			feeds[index].Transaction = tx
			account.NotesCount++
		}

	}

	// Update feed save into database
	if err := global.DB.Scopes(models.FeedTable(models.Feed{
		Feed: types.Feed{
			Platform: account.Platform,
		},
	})).Save(&feeds).Error; err != nil {
		global.Logger.Error("Failed to save updated feeds", feeds)
	}

	// Update account
	global.DB.Save(account)
}

var (
	chFeedOnchain *amqp.Channel
)

func OneFeedOnChain(account *models.Account, feed *models.Feed) (string, string, error) {
	var err error = nil

	// Prepare MQ channel
	if chFeedOnchain == nil {
		chFeedOnchain, err = commonGlobal.MQ.Channel()
		if err != nil {
			global.Logger.Errorf("Failed to prepare MQ channel for onchain with error: %s", err.Error())
			return "", "", err
		}
	}

	// Prepare temp queue
	qTmp, err := chFeedOnchain.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Errorf("Failed to prepare MQ temp queue for account validate with error: %s", err.Error())
		return "", "", err
	}

	// Close tmp queue to prevent memory leak
	defer chFeedOnchain.QueueDelete(qTmp.Name, false, false, false)

	// Prepare respond deliveries
	deliveries, err := chFeedOnchain.Consume(
		qTmp.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Errorf("Failed to register a consumer on account validate response queue with error: %s", err.Error())
		return "", "", err
	}

	var (
		onChainRequestBytes  []byte
		onChainResponseBytes []byte
	)

	if onChainRequestBytes, err = json.Marshal(&commonTypes.OnChainRequest{
		FeedID:               feed.ID,
		CrossbellCharacterID: account.CrossbellCharacterID,
		Platform:             account.Platform,
		Username:             account.Username,
		RawFeed:              feed.RawFeed,
	}); err != nil {
		global.Logger.Errorf("Failed to parse OnChain work for feed %s#%d with error: %s", account.Platform, feed.ID, err.Error())
		utils.AccountOnChainPause(account, fmt.Sprintf("Failed to parse OnChain work for feed %s#%d", account.Platform, feed.ID))
		return "", "", err
	}

	// Prepare corr id
	corrId := uuid.NewString()

	// Prepare ctx
	ctx, cancel := context.WithTimeout(context.Background(), commonConsts.MQSETTINGS_OnChainRequestTimeOut)
	defer cancel()

	// Request validate
	err = chFeedOnchain.PublishWithContext(
		ctx,
		"",
		commonConsts.MQSETTINGS_OnChainRPCQueueName,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/json",
			CorrelationId: corrId,
			ReplyTo:       qTmp.Name,
			Body:          onChainRequestBytes,
		},
	)
	if err != nil {
		global.Logger.Errorf("Failed to dispatch OnChain work for feed %s#%d with error: %s", account.Platform, feed.ID, err.Error())
		utils.AccountOnChainPause(account, fmt.Sprintf("Failed to dispatch OnChain work for feed %s#%d", account.Platform, feed.ID))
		return "", "", err
	}

	// Receive validate response
	global.Logger.Debug("Waiting onchain response from MQ...")
	for d := range deliveries {
		if corrId == d.CorrelationId {
			onChainResponseBytes = d.Body
			break
		}
	}
	// Process response
	global.Logger.Debug("Successfully received onchain response")
	var workRespond commonTypes.OnChainRespond
	if err := json.Unmarshal(onChainResponseBytes, &workRespond); err != nil {
		global.Logger.Errorf("Failed to parse respond: %s", string(onChainResponseBytes))
		utils.AccountOnChainPause(account, fmt.Sprintf("Failed to parse respond: %s", string(onChainResponseBytes)))
		return "", "", err
	}

	// Validate response
	if !workRespond.IsSucceeded {
		global.Logger.Errorf("Failed to finish OnChain work for feed %s#%d with error: %s", account.Platform, feed.ID, workRespond.Message)
		utils.AccountOnChainPause(account, fmt.Sprintf("Failed to finish OnChain work for feed %s#%d", account.Platform, feed.ID))
		return workRespond.IPFSUri, workRespond.Transaction, fmt.Errorf(workRespond.Message)
	}

	return workRespond.IPFSUri, workRespond.Transaction, nil

}
