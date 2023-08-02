package proto

import (
	"github.com/c12s/magnetar/pkg/magnetar"
	"github.com/golang/protobuf/proto"
)

type protoMarshaller struct {
}

func NewMarshaller() magnetar.Marshaller {
	return &protoMarshaller{}
}

func (p protoMarshaller) MarshalRegistrationReq(req magnetar.RegistrationReq) ([]byte, error) {
	protoReq := &RegistrationReq{}
	protoReq, err := protoReq.fromDomain(req)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(protoReq)
}

func (p protoMarshaller) UnmarshalRegistrationReq(reqMarshalled []byte) (*magnetar.RegistrationReq, error) {
	protoReq := RegistrationReq{}
	err := proto.Unmarshal(reqMarshalled, &protoReq)
	if err != nil {
		return nil, err
	}
	return protoReq.toDomain()
}

func (p protoMarshaller) MarshalRegistrationResp(resp magnetar.RegistrationResp) ([]byte, error) {
	protoResp := &RegistrationResp{}
	protoResp, err := protoResp.fromDomain(resp)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(protoResp)
}

func (p protoMarshaller) UnmarshalRegistrationResp(resp []byte) (*magnetar.RegistrationResp, error) {
	protoResp := RegistrationResp{}
	err := proto.Unmarshal(resp, &protoResp)
	if err != nil {
		return nil, err
	}
	return protoResp.toDomain()
}

func (p protoMarshaller) MarshalLabel(label magnetar.Label) ([]byte, error) {
	protoLabel := &Label{}
	protoLabel, err := protoLabel.fromDomain(label)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(protoLabel)
}

func (p protoMarshaller) UnmarshalLabel(labelMarshalled []byte) (magnetar.Label, error) {
	protoLabel := Label{}
	err := proto.Unmarshal(labelMarshalled, &protoLabel)
	if err != nil {
		return nil, err
	}
	return protoLabel.toDomain()
}
