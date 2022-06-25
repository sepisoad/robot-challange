package eventpublisher_test

import (
	"errors"
	"testing"

	. "github.com/sepisoad/robot-challange/shared/nats-mocks/mock"
	"github.com/sepisoad/robot-challange/shared/services/eventpublisher"
	. "github.com/sepisoad/robot-challange/shared/services/robotbroker/mock"
	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

func Test_NewEventPublisherService_Should_Call_CreateNewJetStream(t *testing.T) {
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

	_, err = eventpublisher.NewEventPublisherService(sugarLogger, mockRobotBrokerService)
	g.Expect(err).Should(BeNil())
}

func Test_NewEventPublisherService_Should_Return_Error_If_CreateNewJetStream_Returns_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRobotBrokerService := NewMockRobotBrokerInterface(ctrl)
	mockJetStreamContext := NewMockJetStreamContext(ctrl)

	expectedErr := errors.New(cuid.New())

	mockRobotBrokerService.
		EXPECT().
		CreateNewJetStream().
		Return(mockJetStreamContext, expectedErr)

	g := NewGomegaWithT(t)

	logger, err := zap.NewDevelopment()
	sugarLogger := logger.Sugar()
	g.Expect(err).Should(BeNil())

	_, err = eventpublisher.NewEventPublisherService(sugarLogger, mockRobotBrokerService)
	g.Expect(err).Should(Equal(expectedErr))
}
