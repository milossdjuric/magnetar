package repos

import (
	"errors"
	"github.com/c12s/magnetar/domain"
	"github.com/golang/protobuf/proto"
)

func MarshalNode(node domain.Node) ([]byte, error) {
	protoNode := &NodeDto{}
	protoNode, err := protoNode.fromDomain(node)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(protoNode)
}

func UnmarshalNode(nodeMarshalled []byte) (*domain.Node, error) {
	protoNode := NodeDto{}
	err := proto.Unmarshal(nodeMarshalled, &protoNode)
	if err != nil {
		return nil, err
	}
	return protoNode.toDomain()
}

func MarshalLabel(label domain.Label) ([]byte, error) {
	protoLabel := &LabelDto{}
	protoLabel, err := protoLabel.fromDomain(label)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(protoLabel)
}

func UnmarshalLabel(labelMarshalled []byte) (domain.Label, error) {
	protoLabel := LabelDto{}
	err := proto.Unmarshal(labelMarshalled, &protoLabel)
	if err != nil {
		return nil, err
	}
	return protoLabel.toDomain()
}

func (x *NodeDto) fromDomain(node domain.Node) (*NodeDto, error) {
	var protoLabels []*LabelDto
	for _, label := range node.Labels {
		protoLabel := &LabelDto{}
		protoLabel, err := protoLabel.fromDomain(label)
		if err != nil {
			return nil, err
		}
		protoLabels = append(protoLabels, protoLabel)
	}
	return &NodeDto{
		Id:     node.Id.Value,
		Labels: protoLabels,
	}, nil
}

func (x *NodeDto) toDomain() (*domain.Node, error) {
	var labels []domain.Label
	for _, protoLabel := range x.Labels {
		label, err := protoLabel.toDomain()
		if err != nil {
			return nil, err
		}
		labels = append(labels, label)
	}
	return &domain.Node{
		Id: domain.NodeId{
			Value: x.Id,
		},
		Labels: labels,
	}, nil
}

func (x *LabelDto) fromDomain(label domain.Label) (*LabelDto, error) {
	value := &ValueDto{}
	value, err := value.fromDomain(label.Value())
	if err != nil {
		return nil, err
	}
	return &LabelDto{
		Key:   label.Key(),
		Value: value,
	}, nil
}

func (x *LabelDto) toDomain() (domain.Label, error) {
	var label domain.Label
	var err error
	switch x.Value.Type {
	case ValueDto_Bool:
		protoValue := &BoolValueDto{}
		err = proto.Unmarshal(x.Value.Marshalled, protoValue)
		if err == nil {
			label = domain.NewBoolLabel(x.Key, protoValue.Value)
		}
	case ValueDto_Float64:
		protoValue := &Float64ValueDto{}
		err = proto.Unmarshal(x.Value.Marshalled, protoValue)
		if err == nil {
			label = domain.NewFloat64Label(x.Key, protoValue.Value)
		}
	case ValueDto_String:
		protoValue := &StringValueDto{}
		err = proto.Unmarshal(x.Value.Marshalled, protoValue)
		if err == nil {
			label = domain.NewStringLabel(x.Key, protoValue.Value)
		}
	default:
		err = errors.New("unsupported data type")
	}
	return label, err
}

func (x *ValueDto) fromDomain(value interface{}) (*ValueDto, error) {
	var marshalled []byte
	var valueType ValueDto_ValueTypeDto
	var err error
	switch value.(type) {
	case bool:
		marshalled, err = proto.Marshal(&BoolValueDto{Value: value.(bool)})
		valueType = ValueDto_Bool
	case float64:
		marshalled, err = proto.Marshal(&Float64ValueDto{Value: value.(float64)})
		valueType = ValueDto_Float64
	case string:
		marshalled, err = proto.Marshal(&StringValueDto{Value: value.(string)})
		valueType = ValueDto_String
	default:
		err = errors.New("unsupported data type")
	}
	return &ValueDto{
		Marshalled: marshalled,
		Type:       valueType,
	}, err
}
