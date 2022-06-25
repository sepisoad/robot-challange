package robotbroker

import (
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type robotBrokerService struct {
	logger         *zap.SugaredLogger
	natsConnection *nats.Conn
}

// NewRobotBrokerService creates an concrete instance of RobotBrokerInterface
func NewRobotBrokerService(
	logger *zap.SugaredLogger,
	clientName string,
	natsUrl string) (RobotBrokerInterface, error) {
	robotBrokerService := robotBrokerService{
		logger: logger,
	}

	natsConnection, err := robotBrokerService.createNatsConnection(
		clientName,
		natsUrl)
	if err != nil {
		return nil, err
	}

	robotBrokerService.natsConnection = natsConnection

	err = robotBrokerService.configureStreams()
	if err != nil {
		return nil, err
	}

	return &robotBrokerService, nil
}

// Close closes the connection to the broker
func (s *robotBrokerService) Close() {
	if s.natsConnection != nil {
		s.natsConnection.Close()
		s.natsConnection = nil
	}
}

// CreateNewJetStream creates a message stream which is persisted on disk
func (s *robotBrokerService) CreateNewJetStream() (nats.JetStreamContext, error) {
	return s.natsConnection.JetStream(nats.PublishAsyncMaxPending(256))
}

func (s *robotBrokerService) createNatsConnection(
	clientName string,
	natsUrl string) (*nats.Conn, error) {
	opts := []nats.Option{nats.Name(clientName)}
	opts = s.getNatsConnectionOptions(opts)

	return nats.Connect(natsUrl, opts...)
}

func (s *robotBrokerService) getNatsConnectionOptions(opts []nats.Option) []nats.Option {
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))

	// Never give up
	opts = append(opts, nats.MaxReconnects(-1))

	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		if !nc.IsClosed() {
			s.logger.Warnf("Disconnected due to: %s", err)
		}
	}))

	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		s.logger.Warnf("Reconnected [%s]", nc.ConnectedUrl())
	}))

	return opts
}

func (s *robotBrokerService) configureStreams() (err error) {
	var jetStream nats.JetStreamContext
	if jetStream, err = s.CreateNewJetStream(); err != nil {
		return err
	}

	if _, err := jetStream.AddStream(&nats.StreamConfig{
		Name: "RobotStream",
		Subjects: []string{
			SUBJECT_TASK,
			SUBJECT_ROBOT,
		},
		MaxAge:  time.Hour * 24,
		Storage: nats.FileStorage,
	}); err != nil {
		return err
	}

	return nil
}
