package startup

import (
	"context"
	"errors"
	"github.com/c12s/magnetar/internal/configs"
	"github.com/c12s/magnetar/internal/handlers"
	"github.com/c12s/magnetar/internal/marshallers/proto"
	"github.com/c12s/magnetar/internal/repos"
	"github.com/c12s/magnetar/internal/services"
	"github.com/c12s/magnetar/pkg/api"
	"github.com/c12s/magnetar/pkg/messaging/nats"
	"log"
	"sync"
)

type app struct {
	config                    *configs.Config
	gracefulShutdownProcesses []func(group *sync.WaitGroup)
	shutdownProcesses         []func()
}

func NewAppWithConfig(config *configs.Config) (*app, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}
	return &app{
		config: config,
	}, nil
}

func (a *app) Start() error {
	natsConn, err := NewNatsConn(a.config.NatsAddress())
	if err != nil {
		return err
	}

	a.shutdownProcesses = append(a.shutdownProcesses, func() {
		log.Println("closing nats conn")
		natsConn.Close()
	})

	etcdClient, err := NewEtcdClient(a.config.EtcdAddress())
	if err != nil {
		return err
	}
	a.shutdownProcesses = append(a.shutdownProcesses, func() {
		log.Println("closing etcd client conn")
		err := etcdClient.Close()
		if err != nil {
			log.Println(err)
		}
	})

	nodeRepo, err := repos.NewNodeEtcdRepo(etcdClient, proto.NewProtoNodeMarshaller(), proto.NewProtoLabelMarshaller())
	if err != nil {
		return err
	}

	regReqSubscriber, err := nats.NewSubscriber(natsConn, api.RegistrationSubject, "magnetar")
	if err != nil {
		return err
	}
	a.gracefulShutdownProcesses = append(a.gracefulShutdownProcesses, func(wg *sync.WaitGroup) {
		err := regReqSubscriber.Unsubscribe()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("registration req subscriber gracefully stopped")
		}
		wg.Done()
	})

	regRespPublisher, err := nats.NewPublisher(natsConn)

	registrationService, err := services.NewRegistrationService(nodeRepo)
	if err != nil {
		return err
	}
	registrationHandler, err := handlers.NewAsyncRegistrationHandler(regReqSubscriber, regRespPublisher, *registrationService)
	if err != nil {
		return err
	}
	err = registrationHandler.Handle()
	if err != nil {
		return err
	}

	nodeService, err := services.NewNodeService(nodeRepo)
	if err != nil {
		return err
	}
	labelService, err := services.NewLabelService(nodeRepo)
	if err != nil {
		return err
	}

	magnetarServer, err := handlers.NewMagnetarGrpcServer(*nodeService, *labelService)
	server, err := startServer(a.config.ServerAddress(), magnetarServer)
	if err != nil {
		return err
	}
	a.gracefulShutdownProcesses = append(a.gracefulShutdownProcesses, func(wg *sync.WaitGroup) {
		server.GracefulStop()
		log.Println("grpc server gracefully stopped")
		wg.Done()
	})

	return nil
}

func (a *app) GracefulShutdown(ctx context.Context) {
	// call all shutdown processes after a timeout or graceful shutdown processes completion
	defer a.shutdown()

	// wait for all graceful shutdown processes to complete
	wg := &sync.WaitGroup{}
	wg.Add(len(a.gracefulShutdownProcesses))

	for _, gracefulShutdownProcess := range a.gracefulShutdownProcesses {
		go gracefulShutdownProcess(wg)
	}

	// notify when graceful shutdown processes are done
	gracefulShutdownDone := make(chan struct{})
	go func() {
		wg.Wait()
		gracefulShutdownDone <- struct{}{}
	}()

	// wait for graceful shutdown processes to complete or for ctx timeout
	select {
	case <-ctx.Done():
		log.Println("ctx timeout ... shutting down")
	case <-gracefulShutdownDone:
		log.Println("app gracefully stopped")
	}
}

func (a *app) shutdown() {
	for _, shutdownProcess := range a.shutdownProcesses {
		shutdownProcess()
	}
}
