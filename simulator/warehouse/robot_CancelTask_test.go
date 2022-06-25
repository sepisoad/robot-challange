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

func Test_CancelTask_Should_Cancel_Given_Task(t *testing.T) {
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

	taskId, _, _ := sut.EnqueueTask("E N E N")
	err = sut.CancelTask(taskId)
	g.Expect(err).Should(BeNil())

	robotState := sut.CurrentState()
	g.Expect(robotState.X).Should(Not(Equal(2)))
	g.Expect(robotState.Y).Should(Not(Equal(2)))
}
