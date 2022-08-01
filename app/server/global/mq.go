package global

import "github.com/nats-io/nats.go"

var MQ *nats.Conn
var MQJS nats.JetStreamContext
