package proto

import (
	"github.com/c12s/magnetar/internal/domain"
	"github.com/c12s/magnetar/pkg/api"
)

func GetNodeReqToDomain(req *api.GetNodeReq) (*domain.GetNodeReq, error) {
	return &domain.GetNodeReq{
		Id: domain.NodeId{
			Value: req.NodeId,
		},
	}, nil
}

func GetNodeRespFromDomain(resp domain.GetNodeResp) (*api.GetNodeResp, error) {
	nodeProto, err := NodeStringifiedFromDomain(resp.Node)
	if err != nil {
		return nil, err
	}
	return &api.GetNodeResp{
		Node: nodeProto,
	}, nil
}

func ListNodesReqToDomain(req *api.ListNodesReq) (*domain.ListNodesReq, error) {
	return &domain.ListNodesReq{}, nil
}

func ListNodesRespFromDomain(resp domain.ListNodesResp) (*api.ListNodesResp, error) {
	nodesProto := make([]*api.NodeStringified, len(resp.Nodes))
	for i, node := range resp.Nodes {
		nodeProto, err := NodeStringifiedFromDomain(node)
		if err != nil {
			return nil, err
		}
		nodesProto[i] = nodeProto
	}
	return &api.ListNodesResp{
		Nodes: nodesProto,
	}, nil
}

func QueryNodesReqToDomain(req *api.QueryNodesReq) (*domain.QueryNodesReq, error) {
	selector := make([]domain.Query, 0)
	for _, query := range req.Queries {
		resQuery, err := queryToDomain(query)
		if err != nil {
			return nil, err
		}
		selector = append(selector, *resQuery)
	}
	return &domain.QueryNodesReq{
		Selector: selector,
	}, nil
}

func QueryNodesRespFromDomain(resp domain.QueryNodesResp) (*api.QueryNodesResp, error) {
	protoResp := &api.QueryNodesResp{
		Nodes: make([]*api.NodeStringified, 0),
	}
	for _, node := range resp.Nodes {
		protoNode := &api.NodeStringified{
			Id:     node.Id.Value,
			Labels: make([]*api.LabelStringified, 0),
		}
		for _, label := range node.Labels {
			protoLabel := &api.LabelStringified{
				Key:   label.Key(),
				Value: label.StringValue(),
			}
			protoNode.Labels = append(protoNode.Labels, protoLabel)
		}
		protoResp.Nodes = append(protoResp.Nodes, protoNode)
	}
	return protoResp, nil
}

func PutLabelReqToDomain(req *api.PutLabelReq) (*domain.PutLabelReq, error) {
	label, err := LabelToDomain(req.Label)
	if err != nil {
		return nil, err
	}
	return &domain.PutLabelReq{
		NodeId: domain.NodeId{
			Value: req.NodeId,
		},
		Label: label,
	}, nil
}

func PutLabelRespFromDomain(resp domain.PutLabelResp) (*api.PutLabelResp, error) {
	node, err := NodeStringifiedFromDomain(resp.Node)
	if err != nil {
		return nil, err
	}
	return &api.PutLabelResp{
		Node: node,
	}, nil
}

func DeleteLabelReqToDomain(req *api.DeleteLabelReq) (*domain.DeleteLabelReq, error) {
	return &domain.DeleteLabelReq{
		NodeId: domain.NodeId{
			Value: req.NodeId,
		},
		LabelKey: req.LabelKey,
	}, nil
}

func DeleteLabelRespFromDomain(resp domain.DeleteLabelResp) (*api.DeleteLabelResp, error) {
	node, err := NodeStringifiedFromDomain(resp.Node)
	if err != nil {
		return nil, err
	}
	return &api.DeleteLabelResp{
		Node: node,
	}, nil
}

func queryToDomain(query *api.Query) (*domain.Query, error) {
	shouldBe, err := domain.NewCompResultFromString(query.ShouldBe)
	if err != nil {
		return nil, err
	}
	return &domain.Query{
		LabelKey: query.LabelKey,
		ShouldBe: shouldBe,
		Value:    query.Value,
	}, nil
}
