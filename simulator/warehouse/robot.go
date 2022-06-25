package warehouse

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/sepisoad/robot-challange/shared/services/eventpublisher"
	"github.com/sepisoad/robot-challange/shared/services/idgenerator"
	"go.uber.org/zap"
)

type robot struct {
	logger             *zap.SugaredLogger
	id                 int64
	x                  int
	y                  int
	boardHeight        int
	boardWidth         int
	taskIds            map[int64]bool
	moveMutex          *sync.Mutex
	taskMutex          *sync.Mutex
	idGeneratorService idgenerator.IdGeneratorInterface
}

func NewRobot(
	logger *zap.SugaredLogger,
	id int64,
	x int,
	y int,
	boardHeight int,
	boardWidth int,
	eventpublisherService eventpublisher.EventPublisherInterface,
	idGeneratorService idgenerator.IdGeneratorInterface) (
	RobotInterface,
	error) {

	if err := eventpublisherService.PublishRobotEvent(eventpublisher.RobotEvent{
		EventType: eventpublisher.RobotMoved,
		Id:        id,
		Data: eventpublisher.RobotData{
			X: x,
			Y: y,
		},
	}); err != nil {
		return nil, err
	}

	return &robot{
		logger:             logger,
		id:                 id,
		boardHeight:        boardHeight,
		boardWidth:         boardWidth,
		x:                  x,
		y:                  y,
		taskIds:            make(map[int64]bool),
		moveMutex:          &sync.Mutex{},
		taskMutex:          &sync.Mutex{},
		idGeneratorService: idGeneratorService,
	}, nil
}

func (s *robot) EnqueueTask(commands string) (
	int64,
	chan RobotState,
	chan error) {

	s.taskMutex.Lock()
	taskId := s.idGeneratorService.Generate()
	s.taskIds[taskId] = false
	s.taskMutex.Unlock()

	positionChannel := make(chan RobotState)
	errorChannel := make(chan error)

	go func(
		positionChannel chan RobotState,
		errorChannel chan error) {

		defer close(positionChannel)
		defer close(errorChannel)

		s.moveMutex.Lock()
		defer s.moveMutex.Unlock()

		for _, moveSequenece := range strings.Split(strings.TrimSpace(commands), " ") {
			s.taskMutex.Lock()
			cancelled, found := s.taskIds[taskId]
			s.taskMutex.Unlock()

			if found && cancelled {
				return
			}

			hitTheWall := false

			switch moveSequenece {
			case "S":
				if s.y-1 == -1 {
					hitTheWall = true
				} else {
					s.y = s.y - 1
				}

			case "N":
				if s.y+1 == s.boardHeight {
					hitTheWall = true
				} else {
					s.y = s.y + 1
				}

			case "W":
				if s.x-1 == -1 {
					hitTheWall = true
				} else {
					s.x = s.x - 1
				}

			case "E":
				if s.x+1 == s.boardWidth {
					hitTheWall = true
				} else {
					s.x = s.x + 1
				}
			}

			if hitTheWall {
				errorChannel <- errors.New("robot hit the wall")
			} else {
				positionChannel <- RobotState{
					X: s.x,
					Y: s.y,
				}
			}

			// Simulate moving delay
			time.Sleep(time.Millisecond * 100)
		}
	}(positionChannel, errorChannel)

	return taskId, positionChannel, errorChannel
}

func (s *robot) CancelTask(taskId int64) error {
	s.taskMutex.Lock()
	defer s.taskMutex.Unlock()

	s.taskIds[taskId] = true

	return nil
}

func (s *robot) CurrentState() RobotState {
	return RobotState{
		X: s.x,
		Y: s.y,
	}
}
