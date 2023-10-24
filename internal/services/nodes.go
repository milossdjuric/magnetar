package services

import (
	"github.com/c12s/magnetar/internal/domain"
)

type NodeService struct {
	nodeRepo domain.NodeRepo
}

func NewNodeService(nodeRepo domain.NodeRepo) (*NodeService, error) {
	return &NodeService{
		nodeRepo: nodeRepo,
	}, nil
}

func (n *NodeService) Get(req domain.GetNodeReq) (*domain.GetNodeResp, error) {
	node, err := n.nodeRepo.Get(req.Id)
	if err != nil {
		return nil, err
	}
	return &domain.GetNodeResp{
		Node: *node,
	}, nil
}

func (n *NodeService) List(req domain.ListNodesReq) (*domain.ListNodesResp, error) {
	nodes, err := n.nodeRepo.List()
	if err != nil {
		return nil, err
	}
	return &domain.ListNodesResp{
		Nodes: nodes,
	}, nil
}

func (n *NodeService) Query(req domain.QueryNodesReq) (*domain.QueryNodesResp, error) {
	nodeIds, err := n.nodeRepo.Query(req.Selector)
	if err != nil {
		return nil, err
	}
	nodes := make([]domain.Node, len(nodeIds))
	for i, nodeId := range nodeIds {
		node, err := n.nodeRepo.Get(nodeId)
		if err != nil {
			return nil, err
		}
		nodes[i] = *node
	}
	return &domain.QueryNodesResp{
		Nodes: nodes,
	}, nil
}
