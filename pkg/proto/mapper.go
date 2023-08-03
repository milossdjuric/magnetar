package proto

import (
	"errors"
	"github.com/c12s/magnetar/internal/domain"
	"github.com/c12s/magnetar/pkg/magnetar"
	"github.com/golang/protobuf/proto"
)

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

func (x *GetNodeReq) ToDomain() (*domain.GetNodeReq, error) {
	return &domain.GetNodeReq{
		Id: domain.NodeId{
			Value: x.NodeId,
		},
	}, nil
}

func (x *GetNodeResp) FromDomain(resp domain.GetNodeResp) (*GetNodeResp, error) {
	nodeProto := &NodeStringified{}
	nodeProto, err := nodeProto.fromDomain(resp.Node)
	if err != nil {
		return nil, err
	}
	return &GetNodeResp{
		Node: nodeProto,
	}, nil
}

func (x *ListNodesReq) ToDomain() (*domain.ListNodesReq, error) {
	return &domain.ListNodesReq{}, nil
}

func (x *ListNodesResp) FromDomain(resp domain.ListNodesResp) (*ListNodesResp, error) {
	nodesProto := make([]*NodeStringified, len(resp.Nodes))
	for i, node := range resp.Nodes {
		nodeProto := &NodeStringified{}
		nodeProto, err := nodeProto.fromDomain(node)
		if err != nil {
			return nil, err
		}
		nodesProto[i] = nodeProto
	}
	return &ListNodesResp{
		Nodes: nodesProto,
	}, nil
}

func (x *QueryNodesReq) ToDomain() (*domain.QueryNodesReq, error) {
	selector := make([]magnetar.Query, 0)
	for _, query := range x.Queries {
		resQuery, err := query.toDomain()
		if err != nil {
			return nil, err
		}
		selector = append(selector, *resQuery)
	}
	return &domain.QueryNodesReq{
		Selector: selector,
	}, nil
}

func (x *Query) toDomain() (*magnetar.Query, error) {
	shouldBe, err := magnetar.NewCompResultFromString(x.ShouldBe)
	if err != nil {
		return nil, err
	}
	return &magnetar.Query{
		LabelKey: x.LabelKey,
		ShouldBe: shouldBe,
		Value:    x.Value,
	}, nil
}

func (x *QueryNodesResp) FromDomain(resp domain.QueryNodesResp) (*QueryNodesResp, error) {
	protoResp := &QueryNodesResp{
		Nodes: make([]*NodeStringified, 0),
	}
	for _, node := range resp.Nodes {
		protoNode := &NodeStringified{
			Id:     node.Id.Value,
			Labels: make([]*LabelStringified, 0),
		}
		for _, label := range node.Labels {
			protoLabel := &LabelStringified{
				Key:   label.Key(),
				Value: label.StringValue(),
			}
			protoNode.Labels = append(protoNode.Labels, protoLabel)
		}
		protoResp.Nodes = append(protoResp.Nodes, protoNode)
	}
	return protoResp, nil
}

func (x *PutLabelReq) ToDomain() (*domain.PutLabelReq, error) {
	label, err := x.Label.toDomain()
	if err != nil {
		return nil, err
	}
	return &domain.PutLabelReq{
		NodeId: domain.NodeId{
			Value: x.NodeId,
		},
		Label: label,
	}, nil
}

func (x *PutLabelResp) FromDomain(resp domain.PutLabelResp) (*PutLabelResp, error) {
	node := &NodeStringified{}
	node, err := node.fromDomain(resp.Node)
	if err != nil {
		return nil, err
	}
	return &PutLabelResp{
		Node: node,
	}, nil
}

func (x *DeleteLabelReq) ToDomain() (*domain.DeleteLabelReq, error) {
	return &domain.DeleteLabelReq{
		NodeId: domain.NodeId{
			Value: x.NodeId,
		},
		LabelKey: x.LabelKey,
	}, nil
}

func (x *DeleteLabelResp) FromDomain(resp domain.DeleteLabelResp) (*DeleteLabelResp, error) {
	node := &NodeStringified{}
	node, err := node.fromDomain(resp.Node)
	if err != nil {
		return nil, err
	}
	return &DeleteLabelResp{
		Node: node,
	}, nil
}

func (x *NodeStringified) fromDomain(node domain.Node) (*NodeStringified, error) {
	labels := make([]*LabelStringified, len(node.Labels))
	for i, label := range node.Labels {
		labelProto := &LabelStringified{}
		labelProto, err := labelProto.fromDomain(label)
		if err != nil {
			return nil, err
		}
		labels[i] = labelProto
	}
	return &NodeStringified{
		Id:     node.Id.Value,
		Labels: labels,
	}, nil
}

func (x *LabelStringified) fromDomain(label magnetar.Label) (*LabelStringified, error) {
	return &LabelStringified{
		Key:   label.Key(),
		Value: label.StringValue(),
	}, nil
}

func (x *Node) fromDomain(node domain.Node) (*Node, error) {
	resp := &Node{
		Id:     node.Id.Value,
		Labels: make([]*Label, len(node.Labels)),
	}
	for i, label := range node.Labels {
		protoLabel := &Label{}
		protoLabel, err := protoLabel.fromDomain(label)
		if err != nil {
			return nil, err
		}
		resp.Labels[i] = protoLabel
	}
	return resp, nil
}

func (x *Node) toDomain() (*domain.Node, error) {
	resp := &domain.Node{
		Id: domain.NodeId{
			Value: x.Id,
		},
		Labels: make([]magnetar.Label, len(x.Labels)),
	}
	for i, protoLabel := range x.Labels {
		label, err := protoLabel.toDomain()
		if err != nil {
			return nil, err
		}
		resp.Labels[i] = label
	}
	return resp, nil
}
