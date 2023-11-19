package services

import (
	"context"
	"github.com/c12s/magnetar/internal/domain"
	"log"
)

type NodeService struct {
	nodeRepo domain.NodeRepo
}

func NewNodeService(nodeRepo domain.NodeRepo) (*NodeService, error) {
	return &NodeService{
		nodeRepo: nodeRepo,
	}, nil
}

func (n *NodeService) GetFromNodePool(ctx context.Context, req domain.GetFromNodePoolReq) (*domain.GetFromNodePoolResp, error) {
	node, err := n.nodeRepo.Get(req.Id, "")
	if err != nil {
		return nil, err
	}
	return &domain.GetFromNodePoolResp{
		Node: *node,
	}, nil
}

func (n *NodeService) GetFromOrg(ctx context.Context, req domain.GetFromOrgReq) (*domain.GetFromOrgResp, error) {
	// todo: authorize req
	node, err := n.nodeRepo.Get(req.Id, req.Org)
	if err != nil {
		return nil, err
	}
	return &domain.GetFromOrgResp{
		Node: *node,
	}, nil
}

func (n *NodeService) ClaimOwnership(ctx context.Context, req domain.ClaimOwnershipReq) (*domain.ClaimOwnershipResp, error) {
	// todo: authorize req
	nodes, err := n.nodeRepo.QueryNodePool(req.Query)
	if err != nil {
		return nil, err
	}
	for _, node := range nodes {
		node.Org = req.Org
		err = n.nodeRepo.Put(node)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	// todo: dodaj svaki node u org
	return &domain.ClaimOwnershipResp{
		Nodes: nodes,
	}, nil
}

func (n *NodeService) ListNodePool(ctx context.Context, req domain.ListNodePoolReq) (*domain.ListNodePoolResp, error) {
	nodes, err := n.nodeRepo.ListNodePool()
	if err != nil {
		return nil, err
	}
	return &domain.ListNodePoolResp{
		Nodes: nodes,
	}, nil
}

func (n *NodeService) ListOrgOwnedNodes(ctx context.Context, req domain.ListOrgOwnedNodesReq) (*domain.ListOrgOwnedNodesResp, error) {
	// todo: authorize req
	nodes, err := n.nodeRepo.ListOrgOwnedNodes(req.Org)
	if err != nil {
		return nil, err
	}
	return &domain.ListOrgOwnedNodesResp{
		Nodes: nodes,
	}, nil
}

func (n *NodeService) QueryNodePool(ctx context.Context, req domain.QueryNodePoolReq) (*domain.QueryNodePoolResp, error) {
	nodes, err := n.nodeRepo.QueryNodePool(req.Query)
	if err != nil {
		return nil, err
	}
	return &domain.QueryNodePoolResp{
		Nodes: nodes,
	}, nil
}

func (n *NodeService) QueryOrgOwnedNodes(ctx context.Context, req domain.QueryOrgOwnedNodesReq) (*domain.QueryOrgOwnedNodesResp, error) {
	// todo: authorize req
	nodes, err := n.nodeRepo.QueryOrgOwnedNodes(req.Query, req.Org)
	if err != nil {
		return nil, err
	}
	return &domain.QueryOrgOwnedNodesResp{
		Nodes: nodes,
	}, nil
}
