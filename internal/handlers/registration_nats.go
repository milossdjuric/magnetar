package handlers

import (
	"github.com/c12s/magnetar/internal/domain"
	"github.com/c12s/magnetar/internal/services"
	"github.com/c12s/magnetar/pkg"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type natsRegistrationHandler struct {
	conn                   *nats.Conn
	registrationReqSubject string
	service                services.RegistrationService
	marshaller             pkg.Marshaller
}

func NewNatsRegistrationHandler(conn *nats.Conn, registrationReqSubject string, service services.RegistrationService, marshaller pkg.Marshaller) (domain.RegistrationHandler, error) {
	return natsRegistrationHandler{
		conn:                   conn,
		registrationReqSubject: registrationReqSubject,
		service:                service,
		marshaller:             marshaller,
	}, nil
}

func (n natsRegistrationHandler) Handle() (chan bool, error) {
	subscription, err := n.conn.QueueSubscribe(n.registrationReqSubject, "magnetar", n.handleRegistration)
	if err != nil {
		return nil, err
	}

	subscriptionClosedCh := make(chan bool)
	go func() {
		for subscription.IsValid() {
			time.Sleep(1000 * time.Millisecond)
		}
		subscriptionClosedCh <- true
	}()

	return subscriptionClosedCh, nil
}

func (n natsRegistrationHandler) handleRegistration(msg *nats.Msg) {
	req, err := n.marshaller.UnmarshalRegistrationReq(msg.Data)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := n.service.Register(*req)
	if err != nil {
		log.Println(err)
		return
	}
	marshalledResp, err := n.marshaller.MarshalRegistrationResp(*resp)
	if err != nil {
		log.Println(err)
		return
	}
	err = n.conn.Publish(msg.Reply, marshalledResp)
	if err != nil {
		log.Println(err)
	}
}
