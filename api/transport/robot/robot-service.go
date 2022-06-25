package robot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	robotapiserver "github.com/sepisoad/robot-challange/api-definitions/openapi/robotapi-server"
	"github.com/sepisoad/robot-challange/api/processors"
	"github.com/sepisoad/robot-challange/shared/services/eventpublisher"
	"github.com/sepisoad/robot-challange/shared/services/idgenerator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

type robotService struct {
	logger                           *zap.SugaredLogger
	robotsStatus                     map[int64]processors.RobotStatus
	tasksStatus                      map[int64]processors.TaskStatus
	eventPublisherService            eventpublisher.EventPublisherInterface
	robotStatusMutex                 *sync.Mutex
	taskStatusMutex                  *sync.Mutex
	internalRobotStatusChannels      map[int64]chan map[int64]processors.RobotStatus
	internalRobotStatusChannelsMutex *sync.Mutex
	idGeneratorService               idgenerator.IdGeneratorInterface
}

func NewRobotService(
	logger *zap.SugaredLogger,
	robotStatusChannel chan map[int64]processors.RobotStatus,
	taskStatusChannel chan map[int64]processors.TaskStatus,
	eventPublisherService eventpublisher.EventPublisherInterface,
	idGeneratorService idgenerator.IdGeneratorInterface) (
	robotapiserver.ServerInterface,
	error) {
	service := &robotService{
		logger:                           logger,
		robotsStatus:                     make(map[int64]processors.RobotStatus),
		tasksStatus:                      make(map[int64]processors.TaskStatus),
		eventPublisherService:            eventPublisherService,
		robotStatusMutex:                 &sync.Mutex{},
		taskStatusMutex:                  &sync.Mutex{},
		internalRobotStatusChannels:      make(map[int64]chan map[int64]processors.RobotStatus),
		internalRobotStatusChannelsMutex: &sync.Mutex{},
		idGeneratorService:               idGeneratorService,
	}

	go func(s *robotService, robotStatusChannel chan map[int64]processors.RobotStatus) {
		for robotStatus := range robotStatusChannel {
			s.robotStatusMutex.Lock()
			service.robotsStatus = robotStatus

			s.internalRobotStatusChannelsMutex.Lock()
			for _, internalRobotStatusChannel := range s.internalRobotStatusChannels {
				internalRobotStatusChannel <- s.robotsStatus
			}
			s.internalRobotStatusChannelsMutex.Unlock()

			s.robotStatusMutex.Unlock()
		}
	}(service, robotStatusChannel)

	go func(s *robotService, taskStatusChannel chan map[int64]processors.TaskStatus) {
		for taskStatus := range taskStatusChannel {
			s.taskStatusMutex.Lock()

			for taskId, status := range taskStatus {
				if service.tasksStatus[taskId] != processors.TaskStatusCancelled {
					service.tasksStatus[taskId] = status
				}
			}

			s.taskStatusMutex.Unlock()
		}
	}(service, taskStatusChannel)

	return service, nil
}

// Dashboard returns the web ui
func (s *robotService) Dashboard(ctx echo.Context) error {
	return ctx.File("index.html")
}

// RobotsWebsocket is used by clients to read robots status via websocket
func (s *robotService) RobotsWebsocket(ctx echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		internalRobotStatusChannel := make(chan map[int64]processors.RobotStatus)
		id := s.idGeneratorService.Generate()

		s.internalRobotStatusChannelsMutex.Lock()
		s.internalRobotStatusChannels[id] = internalRobotStatusChannel
		s.internalRobotStatusChannelsMutex.Unlock()

		defer func() {
			s.internalRobotStatusChannelsMutex.Lock()
			delete(s.internalRobotStatusChannels, id)
			close(internalRobotStatusChannel)
			s.internalRobotStatusChannelsMutex.Unlock()
		}()

		robots := s.getRobotStatuses()
		buf, err := json.Marshal(robots)
		if err != nil {
			s.logger.Error(err)

			return
		}

		if err = websocket.Message.Send(ws, string(buf)); err != nil {
			s.logger.Error(err)

			return
		}

		for robotStatus := range internalRobotStatusChannel {
			robots := convertToTransportRobots(robotStatus)

			buf, err := json.Marshal(robots)
			if err != nil {
				s.logger.Error(err)

				break
			}

			if err = websocket.Message.Send(ws, string(buf)); err != nil {
				s.logger.Error(err)

				break
			}

		}
	}).ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

// TasksWebsocket is used by clients to read tasks status via websocket
func (s *robotService) TasksWebsocket(ctx echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			tasksStatus := s.getAllTasks(ctx)
			buf, err := json.Marshal(tasksStatus)
			if err != nil {
				s.logger.Error(err)
				break
			}
			if err = websocket.Message.Send(ws, string(buf)); err != nil {
				s.logger.Error(err)
				break
			}

			time.Sleep(time.Second * 1)
		}
	}).ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

func (s *robotService) GetAllRobots(ctx echo.Context) error {
	robots := s.getRobotStatuses()

	return ctx.JSON(
		http.StatusOK,
		robots)
}

// GetRobot returns a robot by its id
func (s *robotService) GetRobot(ctx echo.Context, robotId robotapiserver.RobotId) error {
	s.robotStatusMutex.Lock()
	defer s.robotStatusMutex.Unlock()

	robot, found := s.robotsStatus[int64(robotId)]
	if !found {
		return getError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("No robot found with Id: %d", robotId))
	}

	return ctx.JSON(
		http.StatusOK,
		robot)
}

// MoveRobot move a robot on grid
func (s *robotService) MoveRobot(ctx echo.Context, robotId robotapiserver.RobotId) error {
	var moveRequest robotapiserver.MoveRobotRequest

	err := ctx.Bind(&moveRequest)
	if err != nil {
		return getError(
			ctx,
			http.StatusBadRequest,
			"Invalid format for MoveRobot request")
	}

	s.robotStatusMutex.Lock()
	defer s.robotStatusMutex.Unlock()

	_, found := s.robotsStatus[int64(robotId)]
	if !found {
		return getError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("No robot found with Id: %d", robotId))
	}

	moveSequeneces := make([]eventpublisher.MoveRobotRequestMoveSequence, 0)

	for _, moveSequence := range moveRequest.MoveSequences {
		moveSequeneces = append(
			moveSequeneces,
			eventpublisher.MoveRobotRequestMoveSequence(moveSequence))
	}

	s.taskStatusMutex.Lock()
	defer s.taskStatusMutex.Unlock()

	taskId := len(s.tasksStatus) + 1
	s.tasksStatus[int64(taskId)] = processors.TaskStatusCreated

	if err = s.eventPublisherService.PublishTaskEvent(eventpublisher.TaskEvent{
		EventType: eventpublisher.TaskCreated,
		Id:        taskId,
		Data: eventpublisher.TaskData{
			RobotId:        int64(robotId),
			MoveSequeneces: moveSequeneces,
		},
	}); err != nil {
		return getError(
			ctx,
			http.StatusInternalServerError,
			err.Error())
	}

	return ctx.JSON(
		http.StatusAccepted,
		robotapiserver.MoveRobotResponse{
			Task: robotapiserver.Task{
				Id: taskId,
			},
		},
	)
}

// MoveRobot returns all tasks with their state
func (s *robotService) GetAllTasks(ctx echo.Context) error {
	tasksStatus := s.getAllTasks(ctx)
	return ctx.JSON(
		http.StatusOK,
		tasksStatus)
}

// MoveRobot returns an specific task by its id
func (s *robotService) GetTask(ctx echo.Context, taskId robotapiserver.TaskId) error {
	s.taskStatusMutex.Lock()
	defer s.taskStatusMutex.Unlock()

	status, found := s.tasksStatus[int64(taskId)]
	if !found {
		return getError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("No task found with Id: %d", taskId))
	}

	return ctx.JSON(
		http.StatusOK,
		robotapiserver.Task{
			Id:     taskId,
			Status: robotapiserver.TaskStatus(status),
		})
}

// MoveRobot cancels a task by its id
func (s *robotService) CancelTask(ctx echo.Context, taskId robotapiserver.TaskId) error {
	s.taskStatusMutex.Lock()
	defer s.taskStatusMutex.Unlock()

	_, found := s.tasksStatus[int64(taskId)]
	if !found {
		return getError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("No task found with Id: %d", taskId))
	}

	if err := s.eventPublisherService.PublishTaskEvent(eventpublisher.TaskEvent{
		EventType: eventpublisher.TaskCancelled,
		Id:        taskId,
	}); err != nil {
		return getError(
			ctx,
			http.StatusInternalServerError,
			err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func getError(ctx echo.Context, code int, message string) error {
	return ctx.JSON(
		code,
		robotapiserver.Error{
			Code:    code,
			Message: message,
		},
	)
}

func (s *robotService) getRobotStatuses() []robotapiserver.Robot {
	s.robotStatusMutex.Lock()
	defer s.robotStatusMutex.Unlock()

	return convertToTransportRobots(s.robotsStatus)
}

func convertToTransportRobots(robotStatus map[int64]processors.RobotStatus) []robotapiserver.Robot {
	robots := make([]robotapiserver.Robot, 0)
	for id, status := range robotStatus {
		robots = append(robots, robotapiserver.Robot{
			Id:        int(id),
			XPosition: status.X,
			YPosition: status.Y,
		})
	}

	sort.SliceStable(robots, func(i, j int) bool {
		return robots[i].Id < robots[j].Id
	})

	return robots
}

func (s *robotService) getAllTasks(ctx echo.Context) []robotapiserver.Task {
	s.taskStatusMutex.Lock()
	defer s.taskStatusMutex.Unlock()

	tasksStatus := make([]robotapiserver.Task, 0)
	for id, status := range s.tasksStatus {
		tasksStatus = append(tasksStatus, robotapiserver.Task{
			Id:     int(id),
			Status: robotapiserver.TaskStatus(status),
		})
	}

	sort.SliceStable(tasksStatus, func(i, j int) bool {
		return tasksStatus[i].Id < tasksStatus[j].Id
	})

	return tasksStatus
}
