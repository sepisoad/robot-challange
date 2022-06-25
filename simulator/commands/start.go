package commands

import (
	"log"

	"github.com/bwmarrin/snowflake"
	"github.com/sepisoad/robot-challange/shared/services/eventpublisher"
	"github.com/sepisoad/robot-challange/shared/services/idgenerator"
	"github.com/sepisoad/robot-challange/shared/services/robotbroker"
	"github.com/sepisoad/robot-challange/simulator/internals/services/config"
	"github.com/sepisoad/robot-challange/simulator/processors"
	"github.com/sepisoad/robot-challange/simulator/warehouse"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type startOptions struct {
	totalRobotNumber int
	boardHeight      int
	boardWidth       int
}

func startCommand() *cobra.Command {
	opt := startOptions{}

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the simualtor",
		Long:  "Start the simualtor",
		Run: func(cmd *cobra.Command, args []string) {
			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			sugarLogger := logger.Sugar()

			configService, err := config.NewConfigService()
			if err != nil {
				sugarLogger.Fatal(err)
			}

			robotBrokerService, err := robotbroker.NewRobotBrokerService(
				sugarLogger,
				"api",
				configService.GetNatsUrl())
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer robotBrokerService.Close()

			eventpublisherService, err := eventpublisher.NewEventPublisherService(
				sugarLogger,
				robotBrokerService)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			snowflakeNode, err := snowflake.NewNode(1)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			idGeneratorService, err := idgenerator.NewIdGeneratorService(snowflakeNode)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			coordinates := getRandomPositions(opt.totalRobotNumber)

			robots := make(map[int64]warehouse.RobotInterface)
			// for idx := 0; idx < opt.totalRobotNumber; idx++ {
			for idx, coord := range coordinates {
				robot, err := warehouse.NewRobot(
					sugarLogger,
					int64(idx),
					coord.x,
					coord.y,
					opt.boardHeight,
					opt.boardWidth,
					eventpublisherService,
					idGeneratorService)
				if err != nil {
					sugarLogger.Fatal(err)
				}

				robots[int64(idx)] = robot
			}

			robotProcessor, err := processors.StartTaskProcessor(
				sugarLogger,
				robotBrokerService,
				robots,
				eventpublisherService)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer robotProcessor.Stop()

			ch := make(chan struct{})

			<-ch
		},
	}

	cmd.Flags().IntVar(&opt.totalRobotNumber, "total-robot-number", 5, "Specify the total number of robots to start simualtion with")
	cmd.Flags().IntVar(&opt.boardHeight, "board-height", 10, "Specify the board height")
	cmd.Flags().IntVar(&opt.boardWidth, "board-width", 10, "Specify the board width")

	return cmd
}

type coordinate struct {
	x int
	y int
}

func getRandomPositions(max int) []coordinate {
	list := make([]coordinate, max)

	for i := 0; i < max; i++ {
		list[i].x = i
		list[i].y = 0
	}

	return list
}
