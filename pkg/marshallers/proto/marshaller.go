package proto

import (
	"errors"
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

func (x *RegistrationReq) fromDomain(req magnetar.RegistrationReq) (*RegistrationReq, error) {
	var protoLabels []*Label
	for _, label := range req.Labels {
		protoLabel := &Label{}
		protoLabel, err := protoLabel.fromDomain(label)
		if err != nil {
			return nil, err
		}
		protoLabels = append(protoLabels, protoLabel)
	}
	return &RegistrationReq{
		Labels: protoLabels,
	}, nil
}

func (x *RegistrationReq) toDomain() (*magnetar.RegistrationReq, error) {
	var labels []magnetar.Label
	for _, protoLabel := range x.Labels {
		label, err := protoLabel.toDomain()
		if err != nil {
			return nil, err
		}
		labels = append(labels, label)
	}
	return &magnetar.RegistrationReq{
		Labels: labels,
	}, nil
}

func (x *RegistrationResp) fromDomain(resp magnetar.RegistrationResp) (*RegistrationResp, error) {
	return &RegistrationResp{
		NodeId: resp.NodeId,
	}, nil
}

func (x *RegistrationResp) toDomain() (*magnetar.RegistrationResp, error) {
	return &magnetar.RegistrationResp{
		NodeId: x.NodeId,
	}, nil
}

func (x *Label) fromDomain(label magnetar.Label) (*Label, error) {
	value := &Value{}
	value, err := value.fromDomain(label.Value())
	if err != nil {
		return nil, err
	}
	return &Label{
		Key:   label.Key(),
		Value: value,
	}, nil
}

func (x *Value) fromDomain(value interface{}) (*Value, error) {
	var marshalled []byte
	var valueType Value_ValueTYpe
	var err error
	switch value.(type) {
	case bool:
		marshalled, err = proto.Marshal(&BoolValue{Value: value.(bool)})
		valueType = Value_Bool
	case float64:
		marshalled, err = proto.Marshal(&Float64Value{Value: value.(float64)})
		valueType = Value_Float64
	case string:
		marshalled, err = proto.Marshal(&StringValue{Value: value.(string)})
		valueType = Value_String
	default:
		err = errors.New("unsupported data type")
	}
	return &Value{
		Marshalled: marshalled,
		Type:       valueType,
	}, err
}

func (x *Label) toDomain() (magnetar.Label, error) {
	var label magnetar.Label
	var err error
	switch x.Value.Type {
	case Value_Bool:
		protoValue := &BoolValue{}
		err = proto.Unmarshal(x.Value.Marshalled, protoValue)
		if err == nil {
			label = magnetar.NewBoolLabel(x.Key, protoValue.Value)
		}
	case Value_Float64:
		protoValue := &Float64Value{}
		err = proto.Unmarshal(x.Value.Marshalled, protoValue)
		if err == nil {
			label = magnetar.NewFloat64Label(x.Key, protoValue.Value)
		}
	case Value_String:
		protoValue := &StringValue{}
		err = proto.Unmarshal(x.Value.Marshalled, protoValue)
		if err == nil {
			label = magnetar.NewStringLabel(x.Key, protoValue.Value)
		}
	default:
		err = errors.New("unsupported data type")
	}
	return label, err
}
