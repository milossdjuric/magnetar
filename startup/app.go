package startup

import (
	"github.com/c12s/magnetar/configs"
	"github.com/c12s/magnetar/handlers"
	"github.com/c12s/magnetar/repos"
	"github.com/c12s/magnetar/services"
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
	nodeRepo, err := repos.NewNodeEtcdRepo(etcdClient)
	if err != nil {
		return err
	}
	registrationService, err := services.NewRegistrationService(nodeRepo)
	if err != nil {
		return err
	}
	registrationHandler, err := handlers.NewNatsRegistrationHandler(natsConn, config.RegistrationSubject(), *registrationService)
	if err != nil {
		return err
	}

	subscriptionClosedCh, err := registrationHandler.Handle()
	if err != nil {
		return err
	}

	<-subscriptionClosedCh

	return nil
}
