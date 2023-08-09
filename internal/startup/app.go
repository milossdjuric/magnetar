package startup

import (
	"github.com/c12s/magnetar/internal/configs"
	"github.com/c12s/magnetar/internal/handlers"
	"github.com/c12s/magnetar/internal/marshallers/proto"
	"github.com/c12s/magnetar/internal/repos"
	"github.com/c12s/magnetar/internal/services"
	"github.com/c12s/magnetar/pkg/api"
	"github.com/c12s/magnetar/pkg/messaging/nats"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func StartApp(config *configs.Config) error {
	natsConn, err := NewNatsConn(config.NatsAddress())
	if err != nil {
		return err
	}
	etcdClient, err := NewEtcdClient(config.EtcdAddress())
	if err != nil {
		return err
	}
	nodeMarshaller := proto.NewProtoNodeMarshaller()
	labelMarshaller := proto.NewProtoLabelMarshaller()
	nodeRepo, err := repos.NewNodeEtcdRepo(etcdClient, nodeMarshaller, labelMarshaller)
	if err != nil {
		return err
	}

	regReqSubscriber, err := nats.NewSubscriber(natsConn, api.RegistrationSubject, "magnetar")
	if err != nil {
		return err
	}
	regRespPublisher, err := nats.NewPublisher(natsConn)

	registrationService, err := services.NewRegistrationService(nodeRepo)
	if err != nil {
		return err
	}
	registrationHandler, err := handlers.NewNatsRegistrationHandler(regReqSubscriber, regRespPublisher, *registrationService)
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
	server, err := startServer(config.ServerAddress(), magnetarServer)
	if err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGINT)
	<-quit

	server.GracefulStop()
	err = regReqSubscriber.Unsubscribe()
	if err != nil {
		log.Println(err)
	}
	natsConn.Close()

	return nil
}
