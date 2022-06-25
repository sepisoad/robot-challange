package eventpublisher

import (
	"encoding/json"

	"github.com/sepisoad/robot-challange/shared/services/robotbroker"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type eventPublisherService struct {
	logger    *zap.SugaredLogger
	jetStream nats.JetStreamContext
}

// NewEventPublisherService creates a concerete instance of EventPublisherInterface
func NewEventPublisherService(
	logger *zap.SugaredLogger,
	robotBrokerService robotbroker.RobotBrokerInterface) (EventPublisherInterface, error) {
	eventPublisherService := eventPublisherService{
		logger: logger,
	}

	jetStream, err := robotBrokerService.CreateNewJetStream()
	if err != nil {
		return nil, err
	}

	eventPublisherService.jetStream = jetStream

	return &eventPublisherService, nil
}

// PublishTaskEvent publishes task event on event queue
func (s *eventPublisherService) PublishTaskEvent(event TaskEvent) error {
	buf, err := json.Marshal(event)
	if err != nil {
		s.logger.Errorf(
			"Failed to serialize TaskEvent message to json. Error: %v", err)

		return err
	}

	if _, err := s.jetStream.Publish(robotbroker.SUBJECT_TASK, buf); err != nil {
		s.logger.Errorf(
			"Failed to publish message to %s. Error: %v",
			robotbroker.SUBJECT_TASK,
			err)

		return err
	}

	return nil
}

// PublishTaskEvent publishes robot event on event queue
func (s *eventPublisherService) PublishRobotEvent(event RobotEvent) error {
	buf, err := json.Marshal(event)
	if err != nil {
		s.logger.Errorf(
			"Failed to serialize RobotEvent message to json. Error: %v", err)

		return err
	}

	if _, err := s.jetStream.Publish(robotbroker.SUBJECT_ROBOT, buf); err != nil {
		s.logger.Errorf(
			"Failed to publish message to %s. Error: %v",
			robotbroker.SUBJECT_ROBOT,
			err)

		return err
	}

	return nil
}
