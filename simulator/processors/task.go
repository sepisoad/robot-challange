package processors

import (
	"encoding/json"
	"sync"

	"github.com/sepisoad/robot-challange/shared/services/eventpublisher"
	"github.com/sepisoad/robot-challange/shared/services/robotbroker"
	"github.com/sepisoad/robot-challange/simulator/warehouse"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type taskMapping struct {
	receivedTaskId int
	robotTaskId    int64
}

type taskProcessor struct {
	logger                  *zap.SugaredLogger
	taskCreatedSubscriber   *nats.Subscription
	taskCancelledSubscriber *nats.Subscription
	robots                  map[int64]warehouse.RobotInterface
	eventpublisherService   eventpublisher.EventPublisherInterface
	taskIdMappings          map[int64][]taskMapping
	taskIdMappingsMutex     *sync.Mutex
}

func StartTaskProcessor(
	logger *zap.SugaredLogger,
	robotBrokerService robotbroker.RobotBrokerInterface,
	robots map[int64]warehouse.RobotInterface,
	eventpublisherService eventpublisher.EventPublisherInterface) (
	processor *taskProcessor,
	err error) {
	var jetStream nats.JetStreamContext

	if jetStream, err = robotBrokerService.CreateNewJetStream(); err != nil {
		return
	}

	taskIdMappings := make(map[int64][]taskMapping)
	// for id := range robots {
	// 	taskIdMappings[id] = make([]taskMapping, 0)
	// }

	processor = &taskProcessor{
		logger:                logger,
		robots:                robots,
		eventpublisherService: eventpublisherService,
		taskIdMappings:        taskIdMappings,
		taskIdMappingsMutex:   &sync.Mutex{},
	}

	if processor.taskCreatedSubscriber, err = jetStream.QueueSubscribe(
		robotbroker.SUBJECT_TASK,
		"simulator-creation-"+robotbroker.SUBJECT_TASK,
		processor.handleTaskCreatedEventRasied); err != nil {
		processor.Stop()

		return
	}

	if processor.taskCreatedSubscriber, err = jetStream.QueueSubscribe(
		robotbroker.SUBJECT_TASK,
		"simulator-cancellation-"+robotbroker.SUBJECT_TASK,
		processor.handleTaskCancelledEventRasied); err != nil {
		processor.Stop()

		return
	}

	return processor, nil
}

func (s *taskProcessor) Stop() {
	if s.taskCreatedSubscriber != nil {
		_ = s.taskCreatedSubscriber.Unsubscribe()
		s.taskCreatedSubscriber = nil
	}

	if s.taskCancelledSubscriber != nil {
		_ = s.taskCancelledSubscriber.Unsubscribe()
		s.taskCancelledSubscriber = nil
	}
}

func (s *taskProcessor) handleTaskCreatedEventRasied(msg *nats.Msg) {
	s.logEnter(msg)

	event := eventpublisher.TaskEvent{}
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		s.logger.Errorf(
			"Failed to de-serialize TaskEvent message. Error: %v",
			err)

		return
	}

	if event.EventType != eventpublisher.TaskCreated {
		return
	}

	robot, found := s.robots[event.Data.RobotId]
	if !found {
		s.logger.Errorf(
			"Robot with Id %d not found",
			event.Data.RobotId)

		return
	}

	commands := ""
	for _, moveSequenece := range event.Data.MoveSequeneces {
		commands = commands + " " + string(moveSequenece)
	}

	taskId, positionChannel, errorChannel := robot.EnqueueTask(commands)

	s.taskIdMappingsMutex.Lock()
	s.taskIdMappings[taskId] = append(
		s.taskIdMappings[taskId],
		taskMapping{
			receivedTaskId: event.Id,
			robotTaskId:    taskId,
		})
	s.taskIdMappingsMutex.Unlock()

	go func(
		taskId int64,
		robotId int64,
		positionChannel chan warehouse.RobotState,
		errorChannel chan error) {
		for {
			positionChannelClosed := false
			errorChannelClosed := false

			select {
			case robotState, ok := <-positionChannel:
				if !ok {
					positionChannelClosed = true

					break
				}

				_ = s.eventpublisherService.PublishRobotEvent(eventpublisher.RobotEvent{
					EventType: eventpublisher.RobotMoved,
					Id:        robotId,
					Data: eventpublisher.RobotData{
						X: robotState.X,
						Y: robotState.Y,
					},
				})

			case err, ok := <-errorChannel:
				if !ok {
					errorChannelClosed = true

					break
				}

				_ = s.eventpublisherService.PublishRobotEvent(eventpublisher.RobotEvent{
					EventType:    eventpublisher.RobotFailedToMove,
					Id:           robotId,
					ErrorMessage: err.Error(),
				})
			}

			_ = errorChannelClosed
			if positionChannelClosed {
				s.eventpublisherService.PublishTaskEvent(eventpublisher.TaskEvent{
					EventType: eventpublisher.TaskCompleted,
					Id:        int(taskId),
				})
				break
			}
		}
	}(taskId, event.Data.RobotId, positionChannel, errorChannel)
}

func (s *taskProcessor) handleTaskCancelledEventRasied(msg *nats.Msg) {
	s.logEnter(msg)

	event := eventpublisher.TaskEvent{}
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		s.logger.Errorf(
			"Failed to de-serialize TaskEvent message. Error: %v",
			err)

		return
	}

	if event.EventType != eventpublisher.TaskCancelled {
		return
	}

	s.taskIdMappingsMutex.Lock()

	for robotId, robot := range s.robots {
		taskMappings := s.taskIdMappings[robotId]

		for _, taskMapping := range taskMappings {
			if taskMapping.receivedTaskId == event.Id {
				go func(robot warehouse.RobotInterface) {
					_ = robot.CancelTask(taskMapping.robotTaskId)
				}(robot)

				break
			}
		}
	}

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
