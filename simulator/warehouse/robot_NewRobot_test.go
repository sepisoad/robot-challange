package warehouse_test

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/sepisoad/robot-challange/shared/services/eventpublisher"
	. "github.com/sepisoad/robot-challange/shared/services/eventpublisher/mock"
	. "github.com/sepisoad/robot-challange/shared/services/idgenerator/mock"
	"github.com/sepisoad/robot-challange/simulator/warehouse"
	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

func Test_NewRobot_Should_Publish_Initial_State_Of_Robot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventpublisherService := NewMockEventPublisherInterface(ctrl)
	mockIdGeneratorService := NewMockIdGeneratorInterface(ctrl)

	g := NewGomegaWithT(t)

	robotId := int64(rand.Intn(1000000))

	mockEventpublisherService.
		EXPECT().
		PublishRobotEvent(gomock.Any()).
		DoAndReturn(
			func(event eventpublisher.RobotEvent) error {

				g.Expect(event.EventType).Should(Equal(eventpublisher.RobotMoved))
				g.Expect(event.Id).Should(Equal(robotId))
				g.Expect(event.ErrorMessage).Should(BeEmpty())
				g.Expect(event.Data.X).Should(BeZero())
				g.Expect(event.Data.Y).Should(BeZero())

				return nil
			})

	logger, err := zap.NewDevelopment()
	sugarLogger := logger.Sugar()
	g.Expect(err).Should(BeNil())

	_, err = warehouse.NewRobot(
		sugarLogger,
		robotId,
		0,
		0,
		10,
		10,
		mockEventpublisherService,
		mockIdGeneratorService)
	g.Expect(err).Should(BeNil())
}

func Test_NewRobot_Should_Return_Error_If_Initial_Publishing_Of_Robot_State_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventpublisherService := NewMockEventPublisherInterface(ctrl)
	mockIdGeneratorService := NewMockIdGeneratorInterface(ctrl)

	g := NewGomegaWithT(t)

	expectedErr := errors.New(cuid.New())

	mockEventpublisherService.
		EXPECT().
		PublishRobotEvent(gomock.Any()).
		Return(expectedErr)

	logger, err := zap.NewDevelopment()
	sugarLogger := logger.Sugar()
	g.Expect(err).Should(BeNil())

	_, err = warehouse.NewRobot(
		sugarLogger,
		0,
		0,
		0,
		10,
		10,
		mockEventpublisherService,
		mockIdGeneratorService)
	g.Expect(err).Should(Equal(expectedErr))
}
