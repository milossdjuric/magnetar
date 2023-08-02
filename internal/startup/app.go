package startup

import (
	"github.com/c12s/magnetar/internal/apis"
	"github.com/c12s/magnetar/internal/configs"
	"github.com/c12s/magnetar/internal/handlers"
	"github.com/c12s/magnetar/internal/repos"
	"github.com/c12s/magnetar/internal/services"
	"github.com/c12s/magnetar/pkg/proto"
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
	marshaller := proto.NewMarshaller()
	nodeRepo, err := repos.NewNodeEtcdRepo(etcdClient, marshaller)
	if err != nil {
		return err
	}

	registrationService, err := services.NewRegistrationService(nodeRepo)
	if err != nil {
		return err
	}
	registrationHandler, err := handlers.NewNatsRegistrationHandler(natsConn, config.RegistrationSubject(), *registrationService, marshaller)
	if err != nil {
		return err
	}
	subscriptionClosedCh, err := registrationHandler.Handle()
	if err != nil {
		return err
	}

	queryService, err := services.NewQueryService(nodeRepo)
	if err != nil {
		return err
	}
	labelService, err := services.NewLabelService(nodeRepo)
	if err != nil {
		return err
	}
	server, err := apis.NewMagnetarGrpcServer(*queryService, *labelService)
	startServer(config.ServerAddress(), server)

	<-subscriptionClosedCh

	return nil
}
