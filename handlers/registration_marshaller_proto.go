package handlers

import (
	"errors"
	"github.com/c12s/magnetar/domain"
	"github.com/golang/protobuf/proto"
)

func UnmarshalReq(reqMarshalled []byte) (*domain.RegistrationReq, error) {
	protoReq := RegistrationReq{}
	err := proto.Unmarshal(reqMarshalled, &protoReq)
	if err != nil {
		return nil, err
	}
	return protoReq.toDomain()
}

func MarshalResp(resp domain.RegistrationResp) ([]byte, error) {
	protoResp := &RegistrationResp{}
	protoResp, err := protoResp.fromDomain(resp)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(protoResp)
}

func (x *RegistrationReq) toDomain() (*domain.RegistrationReq, error) {
	var labels []domain.Label
	for _, protoLabel := range x.Labels {
		label, err := protoLabel.toDomain()
		if err != nil {
			return nil, err
		}
		labels = append(labels, label)
	}
	return &domain.RegistrationReq{
		Labels: labels,
	}, nil
}

func (x *RegistrationResp) fromDomain(resp domain.RegistrationResp) (*RegistrationResp, error) {
	return &RegistrationResp{
		NodeId: resp.NodeId,
	}, nil
}

func (x *Label) toDomain() (domain.Label, error) {
	var label domain.Label
	var err error
	switch x.Value.Type {
	case Value_Bool:
		protoValue := &BoolValue{}
		err = proto.Unmarshal(x.Value.Marshalled, protoValue)
		if err == nil {
			label = domain.NewBoolLabel(x.Key, protoValue.Value)
		}
	case Value_Float64:
		protoValue := &Float64Value{}
		err = proto.Unmarshal(x.Value.Marshalled, protoValue)
		if err == nil {
			label = domain.NewFloat64Label(x.Key, protoValue.Value)
		}
	case Value_String:
		protoValue := &StringValue{}
		err = proto.Unmarshal(x.Value.Marshalled, protoValue)
		if err == nil {
			label = domain.NewStringLabel(x.Key, protoValue.Value)
		}
	default:
		err = errors.New("unsupported data type")
	}
	return label, err
}
