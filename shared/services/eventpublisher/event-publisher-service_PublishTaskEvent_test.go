package eventpublisher_test

import (
	"encoding/json"
	"errors"
	"math/rand"
	"testing"

	. "github.com/sepisoad/robot-challange/shared/nats-mocks/mock"
	"github.com/sepisoad/robot-challange/shared/services/eventpublisher"
	"github.com/sepisoad/robot-challange/shared/services/robotbroker"
	. "github.com/sepisoad/robot-challange/shared/services/robotbroker/mock"
	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"
	"github.com/nats-io/nats.go"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

func Test_PublishTaskEvent_Should_Serialize_Event(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRobotBrokerService := NewMockRobotBrokerInterface(ctrl)
	mockJetStreamContext := NewMockJetStreamContext(ctrl)

	mockRobotBrokerService.
		EXPECT().
		CreateNewJetStream().
		Return(mockJetStreamContext, nil)

	g := NewGomegaWithT(t)

	logger, err := zap.NewDevelopment()
	sugarLogger := logger.Sugar()
	g.Expect(err).Should(BeNil())

	sut, err := eventpublisher.NewEventPublisherService(sugarLogger, mockRobotBrokerService)
	g.Expect(err).Should(BeNil())

	event := eventpublisher.TaskEvent{
		EventType: eventpublisher.TaskCreated,
		Id:        rand.Intn(10000),
		Data: eventpublisher.TaskData{
			RobotId: int64(rand.Intn(10000)),
			MoveSequeneces: []eventpublisher.MoveRobotRequestMoveSequence{
				eventpublisher.EAST,
				eventpublisher.NORTH,
				eventpublisher.EAST,
			},
		},
	}

	mockJetStreamContext.
		EXPECT().
		Publish(robotbroker.SUBJECT_TASK, gomock.Any()).
		DoAndReturn(
			func(_ string, data []byte, _ ...nats.PubOpt) (*nats.PubAck, error) {

				var providedEvent eventpublisher.TaskEvent

				err := json.Unmarshal(data, &providedEvent)
				g.Expect(err).Should(BeNil())
				g.Expect(providedEvent).Should(Equal(event))

				return nil, nil
			})

	err = sut.PublishTaskEvent(event)
	g.Expect(err).Should(BeNil())
}

func Test_PublishTaskEvent_Should_Call_Publish_Method(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRobotBrokerService := NewMockRobotBrokerInterface(ctrl)
	mockJetStreamContext := NewMockJetStreamContext(ctrl)

	mockRobotBrokerService.
		EXPECT().
		CreateNewJetStream().
		Return(mockJetStreamContext, nil)

	g := NewGomegaWithT(t)

	logger, err := zap.NewDevelopment()
	sugarLogger := logger.Sugar()
	g.Expect(err).Should(BeNil())

	sut, err := eventpublisher.NewEventPublisherService(sugarLogger, mockRobotBrokerService)
	g.Expect(err).Should(BeNil())

	mockJetStreamContext.
		EXPECT().
		Publish(robotbroker.SUBJECT_TASK, gomock.Any()).
		Return(nil, nil)

	event := eventpublisher.TaskEvent{}

	err = sut.PublishTaskEvent(event)
	g.Expect(err).Should(BeNil())
}

func Test_PublishTaskEvent_Should_Return_Error_If_Publish_Return_Error(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRobotBrokerService := NewMockRobotBrokerInterface(ctrl)
	mockJetStreamContext := NewMockJetStreamContext(ctrl)

	mockRobotBrokerService.
		EXPECT().
		CreateNewJetStream().
		Return(mockJetStreamContext, nil)

	g := NewGomegaWithT(t)

	logger, err := zap.NewDevelopment()
	sugarLogger := logger.Sugar()
	g.Expect(err).Should(BeNil())

	sut, err := eventpublisher.NewEventPublisherService(sugarLogger, mockRobotBrokerService)
	g.Expect(err).Should(BeNil())

	expectedErr := errors.New(cuid.New())

	mockJetStreamContext.
		EXPECT().
		Publish(robotbroker.SUBJECT_TASK, gomock.Any()).
		Return(nil, expectedErr)

	event := eventpublisher.TaskEvent{}

	err = sut.PublishTaskEvent(event)
	g.Expect(err).Should(Equal(expectedErr))
}
