package handlers

import (
	"github.com/c12s/magnetar/domain"
	"github.com/c12s/magnetar/services"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type natsRegistrationHandler struct {
	conn                   *nats.Conn
	registrationReqSubject string
	service                services.RegistrationService
}

func NewNatsRegistrationHandler(conn *nats.Conn, registrationReqSubject string, service services.RegistrationService) (domain.RegistrationHandler, error) {
	return natsRegistrationHandler{
		conn:                   conn,
		registrationReqSubject: registrationReqSubject,
		service:                service,
	}, nil
}

func (nrh natsRegistrationHandler) Handle() (chan bool, error) {
	subscription, err := nrh.conn.QueueSubscribe(nrh.registrationReqSubject, "magnetar", nrh.handleRegistration)
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

func (nrh natsRegistrationHandler) handleRegistration(msg *nats.Msg) {
	req, err := UnmarshalReq(msg.Data)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := nrh.service.Register(*req)
	if err != nil {
		log.Println(err)
		return
	}
	marshalledResp, err := MarshalResp(*resp)
	if err != nil {
		log.Println(err)
		return
	}
	err = nrh.conn.Publish(msg.Reply, marshalledResp)
	if err != nil {
		log.Println(err)
	}
}
