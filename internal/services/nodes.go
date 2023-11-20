package services

import (
	"context"
	"github.com/c12s/magnetar/internal/domain"
	oortapi "github.com/c12s/oort/pkg/api"
	"log"
)

type NodeService struct {
	nodeRepo      domain.NodeRepo
	evaluator     oortapi.OortEvaluatorClient
	administrator *oortapi.AdministrationAsyncClient
}

func NewNodeService(nodeRepo domain.NodeRepo, evaluator oortapi.OortEvaluatorClient, administrator *oortapi.AdministrationAsyncClient) (*NodeService, error) {
	return &NodeService{
		nodeRepo:      nodeRepo,
		evaluator:     evaluator,
		administrator: administrator,
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
	subject, ok := ctx.Value("subject").(*oortapi.Resource)
	if !ok {
		return nil, domain.ErrForbidden
	}
	authzResp, err := n.evaluator.Authorize(ctx, &oortapi.AuthorizationReq{
		Subject:        subject,
		PermissionName: "node.get",
		Object: &oortapi.Resource{
			Id:   req.Id.Value,
			Kind: "node",
		},
	})
	if err != nil || !authzResp.Authorized {
		log.Println(err)
		return nil, domain.ErrForbidden
	}
	node, err := n.nodeRepo.Get(req.Id, req.Org)
	if err != nil {
		return nil, err
	}
	return &domain.GetFromOrgResp{
		Node: *node,
	}, nil
}

func (n *NodeService) ClaimOwnership(ctx context.Context, req domain.ClaimOwnershipReq) (*domain.ClaimOwnershipResp, error) {
	subject, ok := ctx.Value("subject").(*oortapi.Resource)
	if !ok {
		return nil, domain.ErrForbidden
	}
	authzResp, err := n.evaluator.Authorize(ctx, &oortapi.AuthorizationReq{
		Subject:        subject,
		PermissionName: "node.put",
		Object: &oortapi.Resource{
			Id:   req.Org,
			Kind: "org",
		},
	})
	if err != nil || !authzResp.Authorized {
		log.Println(err)
		return nil, domain.ErrForbidden
	}
	nodes, err := n.nodeRepo.QueryNodePool(req.Query)
	if err != nil {
		return nil, err
	}
	for _, node := range nodes {
		err = n.nodeRepo.Delete(node)
		if err != nil {
			log.Println(err)
			continue
		}
		node.Org = req.Org
		err = n.nodeRepo.Put(node)
		if err != nil {
			log.Println(err)
			continue
		}
		err = n.administrator.SendRequest(&oortapi.CreateInheritanceRelReq{
			From: &oortapi.Resource{
				Id:   req.Org,
				Kind: "org",
			},
			To: &oortapi.Resource{
				Id:   node.Id.Value,
				Kind: "node",
			},
		}, func(resp *oortapi.AdministrationAsyncResp) {
			if resp.Error != "" {
				log.Println(resp.Error)
			}
		})
		if err != nil {
			log.Println(err)
		}
	}
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
	subject, ok := ctx.Value("subject").(*oortapi.Resource)
	if !ok {
		return nil, domain.ErrForbidden
	}
	authzResp, err := n.evaluator.Authorize(ctx, &oortapi.AuthorizationReq{
		Subject:        subject,
		PermissionName: "node.get",
		Object: &oortapi.Resource{
			Id:   req.Org,
			Kind: "org",
		},
	})
	if err != nil || !authzResp.Authorized {
		log.Println(err)
		return nil, domain.ErrForbidden
	}
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
	subject, ok := ctx.Value("subject").(*oortapi.Resource)
	if !ok {
		return nil, domain.ErrForbidden
	}
	authzResp, err := n.evaluator.Authorize(ctx, &oortapi.AuthorizationReq{
		Subject:        subject,
		PermissionName: "node.get",
		Object: &oortapi.Resource{
			Id:   req.Org,
			Kind: "org",
		},
	})
	if err != nil || !authzResp.Authorized {
		log.Println(err)
		return nil, domain.ErrForbidden
	}
	nodes, err := n.nodeRepo.QueryOrgOwnedNodes(req.Query, req.Org)
	if err != nil {
		return nil, err
	}
	return &domain.QueryOrgOwnedNodesResp{
		Nodes: nodes,
	}, nil
}
