package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/nats-io/nats.go"
)

func MQ() error {
	var err error

	global.MQ, err = nats.Connect(config.Config.MQConnString)

	if err != nil {
		return fmt.Errorf("unable to connect to mq: %v", err)
	}

	global.MQJS, err = global.MQ.JetStream()

	if err != nil {
		return fmt.Errorf("failed to create JetStream with error: %v", err)
	}

	return nil
}
