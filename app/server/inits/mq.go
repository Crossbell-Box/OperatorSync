package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/nats-io/nats.go"
)

func MQ() error {
	var err error

	global.MQ, err = nats.Connect(config.Config.MQConnString)
	
	if err != nil {
		return fmt.Errorf("unable to connect to mq: %w", err)
	}

	return nil
}
