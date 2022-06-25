package processors

import (
	"encoding/json"

	"github.com/sepisoad/robot-challange/shared/services/eventpublisher"
	"github.com/sepisoad/robot-challange/shared/services/robotbroker"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

// TaskStatus defines running tasks status
type TaskStatus string

const (
	// TaskStatusCreated is used to denote a Created task
	TaskStatusCreated TaskStatus = "Created"
	// TaskStatusInProgress is used to denote an InProgress task
	TaskStatusInProgress TaskStatus = "InProgress"
	// TaskStatusCompleted is used to denote a Completed task
	TaskStatusCompleted TaskStatus = "Completed"
	// TaskStatusCancelled is used to denote a Cancelled task
	TaskStatusCancelled TaskStatus = "Cancelled"
)

type taskProcessor struct {
	logger            *zap.SugaredLogger
	robotSubscriber   *nats.Subscription
	tasksStatus       map[int64]TaskStatus
	taskStatusChannel chan map[int64]TaskStatus
}

// creates an instance of taskProcessor
func StartTaskProcessor(
	logger *zap.SugaredLogger,
	robotBrokerService robotbroker.RobotBrokerInterface) (
	processor *taskProcessor,
	robotStatusChannel chan map[int64]TaskStatus,
	err error) {
	var jetStream nats.JetStreamContext

	if jetStream, err = robotBrokerService.CreateNewJetStream(); err != nil {
		return
	}

	processor = &taskProcessor{
		logger:            logger,
		tasksStatus:       make(map[int64]TaskStatus),
		taskStatusChannel: make(chan map[int64]TaskStatus),
	}

	if processor.robotSubscriber, err = jetStream.QueueSubscribe(
		robotbroker.SUBJECT_TASK,
		"api-"+robotbroker.SUBJECT_TASK,
		processor.handleTaskEventRaised); err != nil {
		processor.Stop()

		return
	}

	return processor, processor.taskStatusChannel, nil
}

// Stops the the process
func (s *taskProcessor) Stop() {
	if s.robotSubscriber != nil {
		_ = s.robotSubscriber.Unsubscribe()
		s.robotSubscriber = nil
	}

	close(s.taskStatusChannel)
}

func (s *taskProcessor) handleTaskEventRaised(msg *nats.Msg) {
	s.logEnter(msg)

	event := eventpublisher.TaskEvent{}
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		s.logger.Errorf(
			"Failed to de-serialize TaskEvent message. Error: %v",
			err)

		return
	}

	switch event.EventType {
	case eventpublisher.TaskCreated:
		s.tasksStatus[int64(event.Id)] = TaskStatus(event.EventType)
	case eventpublisher.TaskCompleted:
		delete(s.tasksStatus, int64(event.Id))
	}

	s.taskStatusChannel <- s.tasksStatus
}

func (s *taskProcessor) logEnter(msg *nats.Msg) {
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
