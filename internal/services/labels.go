package services

import (
	"context"
	"github.com/c12s/magnetar/internal/domain"
	oortapi "github.com/c12s/oort/pkg/api"
	"log"
)

type LabelService struct {
	nodeRepo  domain.NodeRepo
	evaluator oortapi.OortEvaluatorClient
}

func NewLabelService(nodeRepo domain.NodeRepo, evaluator oortapi.OortEvaluatorClient) (*LabelService, error) {
	return &LabelService{
		nodeRepo:  nodeRepo,
		evaluator: evaluator,
	}, nil
}

func (l *LabelService) PutLabel(ctx context.Context, req domain.PutLabelReq) (*domain.PutLabelResp, error) {
	subject, ok := ctx.Value("subject").(*oortapi.Resource)
	if !ok {
		return nil, domain.ErrForbidden
	}
	authzResp, err := l.evaluator.Authorize(ctx, &oortapi.AuthorizationReq{
		Subject:        subject,
		PermissionName: "node.label.put",
		Object: &oortapi.Resource{
			Id:   req.NodeId.Value,
			Kind: "node",
		},
	})
	if err != nil || !authzResp.Authorized {
		log.Println(err)
		return nil, domain.ErrForbidden
	}
	node, err := l.nodeRepo.Get(req.NodeId, req.Org)
	if err != nil {
		return nil, err
	}
	node, err = l.nodeRepo.PutLabel(*node, req.Label)
	if err != nil {
		return nil, err
	}
	return &domain.PutLabelResp{
		Node: *node,
	}, nil
}

func (l *LabelService) DeleteLabel(ctx context.Context, req domain.DeleteLabelReq) (*domain.DeleteLabelResp, error) {
	subject, ok := ctx.Value("subject").(*oortapi.Resource)
	if !ok {
		return nil, domain.ErrForbidden
	}
	authzResp, err := l.evaluator.Authorize(ctx, &oortapi.AuthorizationReq{
		Subject:        subject,
		PermissionName: "node.label.delete",
		Object: &oortapi.Resource{
			Id:   req.NodeId.Value,
			Kind: "node",
		},
	})
	if err != nil || !authzResp.Authorized {
		log.Println(err)
		return nil, domain.ErrForbidden
	}
	node, err := l.nodeRepo.Get(req.NodeId, req.Org)
	if err != nil {
		return nil, err
	}
	node, err = l.nodeRepo.DeleteLabel(*node, req.LabelKey)
	if err != nil {
		return nil, err
	}
	return &domain.DeleteLabelResp{
		Node: *node,
	}, nil
}
