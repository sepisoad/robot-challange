package robotbroker

import "github.com/nats-io/nats.go"

const (
	CONSUMER_GROUP = "robot-"

	SUBJECT_TASK  = "task"
	SUBJECT_ROBOT = "robot"
)

// RobotBrokerInterface defines contracts for a message broker
type RobotBrokerInterface interface {
	Close()
	CreateNewJetStream() (nats.JetStreamContext, error)
}
