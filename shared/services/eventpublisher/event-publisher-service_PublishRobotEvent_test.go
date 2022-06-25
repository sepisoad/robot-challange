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

func Test_PublishRobotEvent_Should_Serialize_Event(t *testing.T) {
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

	event := eventpublisher.RobotEvent{
		EventType: eventpublisher.RobotMoved,
		Id:        int64(rand.Intn(10000)),
		Data: eventpublisher.RobotData{
			X: rand.Intn(10000),
			Y: rand.Intn(10000),
		},
		ErrorMessage: cuid.New(),
	}

	mockJetStreamContext.
		EXPECT().
		Publish(robotbroker.SUBJECT_ROBOT, gomock.Any()).
		DoAndReturn(
			func(_ string, data []byte, _ ...nats.PubOpt) (*nats.PubAck, error) {

				var providedEvent eventpublisher.RobotEvent

				err := json.Unmarshal(data, &providedEvent)
				g.Expect(err).Should(BeNil())
				g.Expect(providedEvent).Should(Equal(event))

				return nil, nil
			})

	err = sut.PublishRobotEvent(event)
	g.Expect(err).Should(BeNil())
}

func Test_PublishRobotEvent_Should_Call_Publish_Method(t *testing.T) {

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
		Publish(robotbroker.SUBJECT_ROBOT, gomock.Any()).
		Return(nil, nil)

	event := eventpublisher.RobotEvent{}

	err = sut.PublishRobotEvent(event)
	g.Expect(err).Should(BeNil())
}

func Test_PublishRobotEvent_Should_Return_Error_If_Publish_Return_Error(t *testing.T) {

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
		Publish(robotbroker.SUBJECT_ROBOT, gomock.Any()).
		Return(nil, expectedErr)

	event := eventpublisher.RobotEvent{}

	err = sut.PublishRobotEvent(event)
	g.Expect(err).Should(Equal(expectedErr))
}
