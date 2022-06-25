package eventpublisher

// TaskEventType describe task state
type TaskEventType string

const (
	// TaskCreated is used to denote a task event that is Created
	TaskCreated TaskEventType = "Created"
	// TaskCompleted is used to denote a task event that is Completed
	TaskCompleted TaskEventType = "Completed"
	// TaskCancelled is used to denote a task event that is Cancelled
	TaskCancelled TaskEventType = "Cancelled"
)

// TaskEvent describes a task event
type TaskEvent struct {
	EventType TaskEventType `json:"EventType"`
	Id        int           `json:"Id"`
	Data      TaskData      `json:"Data,omitempty"`
}

// MoveRobotRequestMoveSequence describes a movement code
type MoveRobotRequestMoveSequence string

const (
	// EAST denotes a movement code towards EAST
	EAST MoveRobotRequestMoveSequence = "E"
	// NORTH denotes a movement code towards NORTH
	NORTH MoveRobotRequestMoveSequence = "N"
	// SOUTH denotes a movement code towards SOUTH
	SOUTH MoveRobotRequestMoveSequence = "S"
	// WEST denotes a movement code towards WEST
	WEST MoveRobotRequestMoveSequence = "W"
)

// TaskData describes task data
type TaskData struct {
	RobotId        int64                          `json:"RobotId"`
	MoveSequeneces []MoveRobotRequestMoveSequence `json:"MoveSequeneces"`
}

// RobotMovedEventType descries a robot movement event type
type RobotMovedEventType string

const (
	// RobotMoved is used when a RobotMoved
	RobotMoved RobotMovedEventType = "Moved"
	// RobotFailedToMove is used when a RobotFailedToMove
	RobotFailedToMove RobotMovedEventType = "FailedToMove"
)

// RobotEvent describe a RobotEvent
type RobotEvent struct {
	EventType    RobotMovedEventType `json:"EventType"`
	Id           int64               `json:"Id"`
	Data         RobotData           `json:"Data,omitempty"`
	ErrorMessage string              `json:"ErrorMessage,omitempty"`
}

// RobotData show robot's location on grid
type RobotData struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

// EventPublisherInterface defines contract for event publishers
type EventPublisherInterface interface {
	PublishTaskEvent(event TaskEvent) error
	PublishRobotEvent(event RobotEvent) error
}
