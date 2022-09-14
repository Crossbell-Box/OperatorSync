package callback

import (
	"context"
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ValidateHandleFailed(ch *amqp.Channel, d *amqp.Delivery, errCode uint, errMsg string) {
	ValidateHandleResponse(
		ch,
		d,
		false,
		errCode,
		errMsg,
		false,
	)

}

func ValidateHandleResponse(ch *amqp.Channel, d *amqp.Delivery, isSucceeded bool, code uint, msg string, isValid bool) {
	response := commonTypes.ValidateResponse{
		IsSucceeded:            isSucceeded,
		Code:                   code,
		Message:                msg,
		IsValidateStringExists: isValid,
	}

	if responseBytes, err := json.Marshal(&response); err != nil {
		global.Logger.Errorf("Failed to marshall validate respond %v with error: %s", response, err.Error())
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), commonConsts.MQSETTINGS_PublishTimeOut)
		defer cancel()

		err = ch.PublishWithContext(
			ctx,
			"",
			d.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType:   "text/json",
				CorrelationId: d.CorrelationId,
				Body:          responseBytes,
			},
		)

		if err != nil {
			global.Logger.Error("Failed to report validate request with error: ", err.Error())
		}
	}
	d.Ack(false)
}
