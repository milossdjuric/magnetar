package api

import (
	"errors"
	"github.com/c12s/magnetar/pkg/messaging"
	"github.com/golang/protobuf/proto"
	"log"
)

type AsyncRegistrationClient struct {
	publisher         messaging.Publisher
	subscriberFactory func(subject string) messaging.Subscriber
}

func NewAsyncRegistrationClient(publisher messaging.Publisher, subscriberFactory func(subject string) messaging.Subscriber) (*AsyncRegistrationClient, error) {
	if publisher == nil {
		return nil, errors.New("publisher is nil")
	}
	return &AsyncRegistrationClient{
		publisher:         publisher,
		subscriberFactory: subscriberFactory,
	}, nil
}

func (n *AsyncRegistrationClient) Register(req *RegistrationReq, callback RegistrationCallback) error {
	reqMarshalled, err := req.Marshal()
	if err != nil {
		return err
	}

	replySubject := n.publisher.GenerateReplySubject()
	subscriber := n.subscriberFactory(replySubject)
	err = subscriber.Subscribe(func(msg []byte, _ string) {
		resp := &RegistrationResp{}
		err := resp.Unmarshal(msg)
		if err != nil {
			log.Println(err)
			return
		}
		callback(resp)
	})
	if err != nil {
		return err
	}

	// send request
	err = n.publisher.Request(reqMarshalled, RegistrationSubject, replySubject)
	if err != nil {
		_ = subscriber.Unsubscribe()
		return err
	}
	return nil
}

type RegistrationCallback func(resp *RegistrationResp)

type RegistrationReqBuilder struct {
	req *RegistrationReq
}

func NewRegistrationReqBuilder() RegistrationReqBuilder {
	return RegistrationReqBuilder{
		req: &RegistrationReq{
			Labels: make([]*Label, 0),
		},
	}
}

func (r RegistrationReqBuilder) AddBoolLabel(key string, value bool) RegistrationReqBuilder {
	valueMarshalled, err := proto.Marshal(&BoolValue{Value: value})
	if err != nil {
		return r
	}
	return r.addLabel(key, Value_Bool, valueMarshalled)
}

func (r RegistrationReqBuilder) AddFloat64Label(key string, value float64) RegistrationReqBuilder {
	valueMarshalled, err := proto.Marshal(&Float64Value{Value: value})
	if err != nil {
		return r
	}
	return r.addLabel(key, Value_Float64, valueMarshalled)
}

func (r RegistrationReqBuilder) AddStringLabel(key string, value string) RegistrationReqBuilder {
	valueMarshalled, err := proto.Marshal(&StringValue{Value: value})
	if err != nil {
		return r
	}
	return r.addLabel(key, Value_String, valueMarshalled)
}

func (r RegistrationReqBuilder) addLabel(key string, valueType Value_ValueTYpe, valueMarshalled []byte) RegistrationReqBuilder {
	label := &Label{
		Key: key,
		Value: &Value{
			Type:       valueType,
			Marshalled: valueMarshalled,
		},
	}
	r.req.Labels = append(r.req.Labels, label)
	return r
}

func (r RegistrationReqBuilder) Request() *RegistrationReq {
	return r.req
}

func (r RegistrationReqBuilder) Clear() RegistrationReqBuilder {
	return NewRegistrationReqBuilder()
}
