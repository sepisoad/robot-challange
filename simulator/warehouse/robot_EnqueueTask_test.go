package warehouse_test

import (
	"math/rand"
	"testing"

	. "github.com/sepisoad/robot-challange/shared/services/eventpublisher/mock"
	. "github.com/sepisoad/robot-challange/shared/services/idgenerator/mock"
	"github.com/sepisoad/robot-challange/simulator/warehouse"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

func Test_EnqueueTask_Should_Return_Valid_Task_Id(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventpublisherService := NewMockEventPublisherInterface(ctrl)
	mockIdGeneratorService := NewMockIdGeneratorInterface(ctrl)

	mockEventpublisherService.
		EXPECT().
		PublishRobotEvent(gomock.Any()).
		Return(nil)

	var expectedTaskId int64 = int64(rand.Intn(10000))

	mockIdGeneratorService.
		EXPECT().
		Generate().
		Return(expectedTaskId)

	g := NewGomegaWithT(t)

	logger, err := zap.NewDevelopment()
	sugarLogger := logger.Sugar()
	g.Expect(err).Should(BeNil())

	sut, err := warehouse.NewRobot(
		sugarLogger,
		0,
		0,
		0,
		10,
		10,
		mockEventpublisherService,
		mockIdGeneratorService)
	g.Expect(err).Should(BeNil())

	taskId, _, _ := sut.EnqueueTask("N")
	g.Expect(taskId).Should(Equal(expectedTaskId))
}

func Test_EnqueueTask_Should_Send_Updated_Position_Through_Returned_Poisition_Channel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventpublisherService := NewMockEventPublisherInterface(ctrl)
	mockIdGeneratorService := NewMockIdGeneratorInterface(ctrl)

	mockEventpublisherService.
		EXPECT().
		PublishRobotEvent(gomock.Any()).
		Return(nil)

	var expectedTaskId int64 = int64(rand.Intn(10000))

	mockIdGeneratorService.
		EXPECT().
		Generate().
		Return(expectedTaskId)

	g := NewGomegaWithT(t)

	logger, err := zap.NewDevelopment()
	sugarLogger := logger.Sugar()
	g.Expect(err).Should(BeNil())

	sut, err := warehouse.NewRobot(
		sugarLogger,
		0,
		0,
		0,
		10,
		10,
		mockEventpublisherService,
		mockIdGeneratorService)
	g.Expect(err).Should(BeNil())

	_, positionChannel, _ := sut.EnqueueTask("E N")
	g.Expect(positionChannel).Should(Not(BeNil()))

	robotState := <-positionChannel
	g.Expect(robotState.X).Should(Equal(1))
	g.Expect(robotState.Y).Should(Equal(0))

	robotState = <-positionChannel
	g.Expect(robotState.X).Should(Equal(1))
	g.Expect(robotState.Y).Should(Equal(1))
}

func Test_EnqueueTask_Should_Send_Error_Through_Returned_Error_Channel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventpublisherService := NewMockEventPublisherInterface(ctrl)
	mockIdGeneratorService := NewMockIdGeneratorInterface(ctrl)

	mockEventpublisherService.
		EXPECT().
		PublishRobotEvent(gomock.Any()).
		Return(nil)

	var expectedTaskId int64 = int64(rand.Intn(10000))

	mockIdGeneratorService.
		EXPECT().
		Generate().
		Return(expectedTaskId)

	g := NewGomegaWithT(t)

	logger, err := zap.NewDevelopment()
	sugarLogger := logger.Sugar()
	g.Expect(err).Should(BeNil())

	sut, err := warehouse.NewRobot(
		sugarLogger,
		0,
		0,
		0,
		10,
		10,
		mockEventpublisherService,
		mockIdGeneratorService)
	g.Expect(err).Should(BeNil())

	_, positionChannel, errorChannel := sut.EnqueueTask("E N W W")
	g.Expect(positionChannel).Should(Not(BeNil()))

	robotState := <-positionChannel
	g.Expect(robotState.X).Should(Equal(1))
	g.Expect(robotState.Y).Should(Equal(0))

	robotState = <-positionChannel
	g.Expect(robotState.X).Should(Equal(1))
	g.Expect(robotState.Y).Should(Equal(1))

	robotState = <-positionChannel
	g.Expect(robotState.X).Should(Equal(0))
	g.Expect(robotState.Y).Should(Equal(1))

	err = <-errorChannel
	g.Expect(err).Should(Not(BeNil()))
}

func Test_EnqueueTask_Should_Carry_On_Moving_After_Hit_The_Wall(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventpublisherService := NewMockEventPublisherInterface(ctrl)
	mockIdGeneratorService := NewMockIdGeneratorInterface(ctrl)

	mockEventpublisherService.
		EXPECT().
		PublishRobotEvent(gomock.Any()).
		Return(nil)

	var expectedTaskId int64 = int64(rand.Intn(10000))

	mockIdGeneratorService.
		EXPECT().
		Generate().
		Return(expectedTaskId)

	g := NewGomegaWithT(t)

	logger, err := zap.NewDevelopment()
	sugarLogger := logger.Sugar()
	g.Expect(err).Should(BeNil())

	sut, err := warehouse.NewRobot(
		sugarLogger,
		0,
		0,
		0,
		10,
		10,
		mockEventpublisherService,
		mockIdGeneratorService)
	g.Expect(err).Should(BeNil())

	_, positionChannel, errorChannel := sut.EnqueueTask("E N W W E")
	g.Expect(positionChannel).Should(Not(BeNil()))

	robotState := <-positionChannel
	g.Expect(robotState.X).Should(Equal(1))
	g.Expect(robotState.Y).Should(Equal(0))

	robotState = <-positionChannel
	g.Expect(robotState.X).Should(Equal(1))
	g.Expect(robotState.Y).Should(Equal(1))

	robotState = <-positionChannel
	g.Expect(robotState.X).Should(Equal(0))
	g.Expect(robotState.Y).Should(Equal(1))

	err = <-errorChannel
	g.Expect(err).Should(Not(BeNil()))

	robotState = <-positionChannel
	g.Expect(robotState.X).Should(Equal(1))
	g.Expect(robotState.Y).Should(Equal(1))
}
