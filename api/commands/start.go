package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/snowflake"
	robotapiserver "github.com/sepisoad/robot-challange/api-definitions/openapi/robotapi-server"
	"github.com/sepisoad/robot-challange/api/internals/services/config"
	"github.com/sepisoad/robot-challange/api/processors"
	"github.com/sepisoad/robot-challange/api/transport/robot"
	"github.com/sepisoad/robot-challange/shared/services/eventpublisher"
	"github.com/sepisoad/robot-challange/shared/services/idgenerator"
	"github.com/sepisoad/robot-challange/shared/services/robotbroker"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the API",
		Long:  "Start the API",
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

			eventPublisherService, err := eventpublisher.NewEventPublisherService(
				sugarLogger,
				robotBrokerService)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			robotProcessor, robotStatusChannel, err := processors.StartRobotProcessor(
				sugarLogger,
				robotBrokerService)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer robotProcessor.Stop()

			taskProcessor, taskStatusChannel, err := processors.StartTaskProcessor(
				sugarLogger,
				robotBrokerService)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer taskProcessor.Stop()

			swaggerSpec, err := robotapiserver.GetSwagger()
			if err != nil {
				sugarLogger.Fatalf("Error loading swagger spec\n: %v", err)
			}

			swaggerSpec.Servers = nil //TODO:sepi

			snowflakeNode, err := snowflake.NewNode(1)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			idGeneratorService, err := idgenerator.NewIdGeneratorService(snowflakeNode)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			robotService, err := robot.NewRobotService(
				sugarLogger,
				robotStatusChannel,
				taskStatusChannel,
				eventPublisherService,
				idGeneratorService)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			e := echo.New()
			e.Use(echomiddleware.CORSWithConfig(
				echomiddleware.CORSConfig{
					AllowOrigins: []string{"*"},
					AllowHeaders: []string{
						echo.HeaderOrigin,
						echo.HeaderContentType,
						echo.HeaderAccept},
				}))
			e.Use(echomiddleware.Logger()) //TODO:sepi
			e.Use(middleware.OapiRequestValidator(swaggerSpec))

			robotapiserver.RegisterHandlers(e, robotService)

			e.File("/", "index.html")

			e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", configService.GetListeningPort())))
		},
	}

	return cmd
}
