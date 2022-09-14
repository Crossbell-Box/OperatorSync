package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

var (
	chValidateAccount *amqp.Channel
)

func ValidateAccount(crossbellCharacterID string, platform string, username string) (bool, error) {

	var err error = nil

	// Prepare MQ channel
	if chValidateAccount == nil {
		chValidateAccount, err = commonGlobal.MQ.Channel()
		if err != nil {
			global.Logger.Errorf("Failed to prepare MQ channel for account validate with error: %s", err.Error())
			return false, err
		}
	}

	// Prepare temp queue
	qTmp, err := chValidateAccount.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Errorf("Failed to prepare MQ temp queue for account validate with error: %s", err.Error())
		return false, err
	}

	// Close tmp queue to prevent memory leak
	defer chValidateAccount.QueueDelete(qTmp.Name, false, false, false)

	// Prepare respond deliveries
	deliveries, err := chValidateAccount.Consume(
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
		return false, err
	}

	// Prepare validate request data
	var (
		validateRequestBytes  []byte
		validateResponseBytes []byte
	)

	if validateRequestBytes, err = json.Marshal(&commonTypes.ValidateRequest{
		Platform:             platform,
		Username:             username,
		CrossbellCharacterID: crossbellCharacterID,
	}); err != nil {
		global.Logger.Error("Failed to parse validate request to bytes: ", err.Error())
		return false, err
	}

	// Prepare corr id
	corrId := uuid.NewString()

	// Prepare ctx
	ctx, cancel := context.WithTimeout(context.Background(), commonConsts.MQSETTINGS_PublishTimeOut)
	defer cancel()

	// Request validate
	err = chValidateAccount.PublishWithContext(
		ctx,
		"",
		commonConsts.MQSETTINGS_ValidateRPCQueueName,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/json",
			CorrelationId: corrId,
			ReplyTo:       qTmp.Name,
			Body:          validateRequestBytes,
		},
	)
	if err != nil {
		global.Logger.Error("Failed to publish validate request", err.Error())
		return false, err
	}

	// Receive validate response
	global.Logger.Debug("Waiting validate response from MQ...")

	// Set timeout
	timeout := time.After(commonConsts.MQSETTINGS_ValidateRequestTimeOut)
	select {
	case <-timeout:
		// Timeout
		global.Logger.Errorf("Validate request timeout...")
		return false, fmt.Errorf("validate request timeout")

	case d := <-deliveries:
		if corrId == d.CorrelationId {
			validateResponseBytes = d.Body
			break
		}
	}

	// Process response
	global.Logger.Debug("Successfully received validate response")
	var validateResponse commonTypes.ValidateResponse
	if err := json.Unmarshal(validateResponseBytes, &validateResponse); err != nil {
		global.Logger.Error("Failed to parse validate response", err.Error())
		return false, err
	}

	if !validateResponse.IsSucceeded {
		// Something is wrong
		global.Logger.Error("Validate work failed with code ", validateResponse.Code, " and message: ", validateResponse.Message)
		return false, fmt.Errorf(validateResponse.Message)
	}

	return validateResponse.IsValidateStringExists, nil
}
