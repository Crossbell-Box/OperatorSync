package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/nats-io/nats.go"
)

func MQ() error {
	// Note: NATS would drop works if there's no consumer to accept, but it's OK for now.
	// If there's further need for those, using JetStream mode might be better.

	var err error

	global.MQ, err = nats.Connect(config.Config.MQConnString)

	if err != nil {
		return fmt.Errorf("unable to connect to mq: %v", err)
	}

	return nil
}
