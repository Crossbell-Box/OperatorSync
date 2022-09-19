package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/common/global"
	amqp "github.com/rabbitmq/amqp091-go"
)

func MQ(MQConnString string) error {

	// Prepare connection

	var err error

	global.MQ, err = amqp.Dial(MQConnString)
	if err != nil {
		return fmt.Errorf("unable to connect to mq: %v", err)
	}
	//defer conn.Close()

	return nil
}
