package inits

import (
	"github.com/Crossbell-Box/OperatorSync/common/consts"
	"github.com/Crossbell-Box/OperatorSync/common/global"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

func MQ(MQConnString string, logger *zap.SugaredLogger) error {

	global.MQ = nil

	go func() {
		// Listen to connection close event
		for {
			notifyClose := make(chan *amqp.Error)
			// Connect
			global.MQ = nil
			connCounter := 0
			for connCounter < consts.MQSETTINGS_ReconnectLimit {
				connCounter++
				logger.Infof("MQ connecting... (%d/%d)", connCounter, consts.MQSETTINGS_ReconnectLimit)
				conn, err := amqp.Dial(MQConnString)
				if err == nil {
					// Valid connection
					conn.NotifyClose(notifyClose)
					global.MQ = conn
					break
				} else {
					logger.Errorf("Failed to connect to MQ with error: %s", err.Error())
				}
				time.Sleep(consts.MQSETTINGS_ReconnectDelay)
			}

			if connCounter >= consts.MQSETTINGS_ReconnectLimit {
				logger.Fatalf("Failed to connect to MQ, service stopping...")
			} else {
				logger.Infof("MQ connect succeed.")
			}

			// Waiting for error
			for {
				err := <-notifyClose
				if err != nil {
					logger.Errorf("MQ connection closed with error %d (%s), preparing to reconnect", err.Code, err.Error())
					break
				}

			}
		}
	}()

	// Waiting for connection
	for {
		time.Sleep(consts.MQSETTINGS_ReconnectDelay)
		if global.MQ != nil {
			break
		}
	}

	return nil
}
