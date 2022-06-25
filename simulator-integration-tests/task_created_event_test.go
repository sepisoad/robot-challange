package simulator_integration_tests

import (
	"encoding/json"
	"testing"

	"github.com/sepisoad/robot-challange/shared/services/eventpublisher"
	"github.com/sepisoad/robot-challange/shared/services/robotbroker"
	"github.com/sepisoad/robot-challange/simulator-integration-tests/internals/services/config"
	"github.com/nats-io/nats.go"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var internalChannel = make(chan eventpublisher.RobotEvent)
var taskId = 1
var robotId = int64(0)

func Test_Should_Process_Task_Created_Event(t *testing.T) {
	g := NewGomegaWithT(t)

	logger, err := zap.NewDevelopment()
	g.Expect(err).Should(BeNil())

	sugarLogger := logger.Sugar()

	configService, err := config.NewConfigService()
	g.Expect(err).Should(BeNil())

	robotBrokerService, err := robotbroker.NewRobotBrokerService(
		sugarLogger,
		"simulator-integration-tests",
		configService.GetNatsUrl())
	g.Expect(err).Should(BeNil())

	defer robotBrokerService.Close()

	eventpublisherService, err := eventpublisher.NewEventPublisherService(
		sugarLogger,
		robotBrokerService)
	g.Expect(err).Should(BeNil())

	var jetStream nats.JetStreamContext

	jetStream, err = robotBrokerService.CreateNewJetStream()
	g.Expect(err).Should(BeNil())

	subscriber, err := jetStream.QueueSubscribe(
		robotbroker.SUBJECT_ROBOT,
		"simulator-integration-tests-"+robotbroker.SUBJECT_ROBOT,
		handleRobotEventRasied)
	g.Expect(err).Should(BeNil())
	defer subscriber.Unsubscribe()

	err = eventpublisherService.PublishTaskEvent(eventpublisher.TaskEvent{
		EventType: eventpublisher.TaskCreated,
		Id:        taskId,
		Data: eventpublisher.TaskData{
			RobotId: robotId,
			MoveSequeneces: []eventpublisher.MoveRobotRequestMoveSequence{
				eventpublisher.EAST,
				eventpublisher.NORTH,
			},
		},
	})
	g.Expect(err).Should(BeNil())

	// The initial state that is broadcasted
	event := <-internalChannel
	g.Expect(event.EventType).Should(Equal(eventpublisher.RobotMoved))
	g.Expect(event.Data.X).Should(Equal(0))
	g.Expect(event.Data.Y).Should(Equal(0))

	event = <-internalChannel
	g.Expect(event.EventType).Should(Equal(eventpublisher.RobotMoved))
	g.Expect(event.Data.X).Should(Equal(1))
	g.Expect(event.Data.Y).Should(Equal(0))

	event = <-internalChannel
	g.Expect(event.EventType).Should(Equal(eventpublisher.RobotMoved))
	g.Expect(event.Data.X).Should(Equal(1))
	g.Expect(event.Data.Y).Should(Equal(1))
}

func handleRobotEventRasied(msg *nats.Msg) {
	event := eventpublisher.RobotEvent{}
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		return
	}

	if event.EventType != eventpublisher.RobotMoved {
		return
	}

	if event.Id != robotId {
		return
	}

	internalChannel <- event
}
