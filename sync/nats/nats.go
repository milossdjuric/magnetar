package nats

import (
	mPb "github.com/c12s/scheme/magnetar"
	"github.com/golang/protobuf/proto"
	nats "github.com/nats-io/nats.go"
)

type NatsSync struct {
	nc    *nats.Conn
	topic string
}

func New(address, topic string) (*NatsSync, error) {
	nc, err := nats.Connect(address)
	if err != nil {
		return nil, err
	}

	return &NatsSync{
		nc:    nc,
		topic: topic,
	}, nil
}

func (ns *NatsSync) Sub(f func(u *mPb.EventMsg)) {
	ns.nc.QueueSubscribe(ns.topic, "magnetar-service", func(msg *nats.Msg) {
		data := &mPb.EventMsg{}
		err := proto.Unmarshal(msg.Data, data)
		if err != nil {
			f(nil)
		}
		f(data)
	})
	ns.nc.Flush()
}
