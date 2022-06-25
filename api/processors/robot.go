package processors

import (
	"encoding/json"

	"github.com/sepisoad/robot-challange/shared/services/eventpublisher"
	"github.com/sepisoad/robot-challange/shared/services/robotbroker"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

// RobotStatus defines robot status
type RobotStatus struct {
	X int
	Y int
}

type robotProcessor struct {
	logger             *zap.SugaredLogger
	robotSubscriber    *nats.Subscription
	robotsStatus       map[int64]RobotStatus
	robotStatusChannel chan map[int64]RobotStatus
}

// creates an instance of robotProcessor and starts it
func StartRobotProcessor(
	logger *zap.SugaredLogger,
	robotBrokerService robotbroker.RobotBrokerInterface) (
	processor *robotProcessor,
	robotStatusChannel chan map[int64]RobotStatus,
	err error) {
	var jetStream nats.JetStreamContext

	if jetStream, err = robotBrokerService.CreateNewJetStream(); err != nil {
		return
	}

	processor = &robotProcessor{
		logger:             logger,
		robotsStatus:       make(map[int64]RobotStatus),
		robotStatusChannel: make(chan map[int64]RobotStatus),
	}

	if processor.robotSubscriber, err = jetStream.QueueSubscribe(
		robotbroker.SUBJECT_ROBOT,
		"api-move-"+robotbroker.SUBJECT_ROBOT,
		processor.handleRobotMovedEventRaised); err != nil {
		processor.Stop()

		return
	}

	return processor, processor.robotStatusChannel, nil
}

func (s *robotProcessor) Stop() {
	if s.robotSubscriber != nil {
		_ = s.robotSubscriber.Unsubscribe()
		s.robotSubscriber = nil
	}

	close(s.robotStatusChannel)
}

func (s *robotProcessor) handleRobotMovedEventRaised(msg *nats.Msg) {
	s.logEnter(msg)

	event := eventpublisher.RobotEvent{}
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		s.logger.Errorf(
			"Failed to de-serialize RobotEvent message. Error: %v",
			err)

		return
	}

	if event.EventType != eventpublisher.RobotMoved {
		return
	}

	s.robotsStatus[event.Id] = RobotStatus{
		X: event.Data.X,
		Y: event.Data.Y,
	}

	s.robotStatusChannel <- s.robotsStatus
}

func (s *robotProcessor) logEnter(msg *nats.Msg) {
	metadata, err := msg.Metadata()
	if err != nil {
		s.logger.Infof("Received message from subject: %s", msg.Subject)

		return
	}

	s.logger.Infof(
		"Stream: %v, Sequence: %v. Received message from subject: %s",
		metadata.Sequence.Stream,
		metadata.Sequence.Consumer,
		msg.Subject)
}
