package warehouse

type RobotState struct {
	X        int
	Y        int
	HasCrate bool
}

type RobotInterface interface {
	EnqueueTask(commands string) (
		taskId int64,
		positionChannel chan RobotState,
		errorChannel chan error)
	CancelTask(taskId int64) error
	CurrentState() RobotState
}
