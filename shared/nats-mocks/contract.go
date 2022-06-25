package nats

import "github.com/nats-io/nats.go"

type JetStreamContext interface {
	nats.JetStreamContext
}
