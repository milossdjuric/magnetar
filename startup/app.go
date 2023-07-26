package startup

import (
	"github.com/c12s/magnetar/configs"
	"github.com/c12s/magnetar/handlers"
	"github.com/c12s/magnetar/services"
)

func StartApp(config *configs.Config) error {
	natsConn, err := NewNatsConn(config.NatsAddress())
	if err != nil {
		return err
	}
	registrationService := services.NewRegistrationService()
	registrationHandler := handlers.NewNatsRegistrationHandler(natsConn, config.RegistrationSubject(), *registrationService)

	subscriptionClosedCh, err := registrationHandler.Handle()
	if err != nil {
		return err
	}

	<-subscriptionClosedCh

	return nil
}
