package services

import (
	"github.com/c12s/magnetar/internal/domain"
)

type LabelService struct {
	nodeRepo domain.NodeRepo
}

func NewLabelService(nodeRepo domain.NodeRepo) (*LabelService, error) {
	return &LabelService{
		nodeRepo: nodeRepo,
	}, nil
}

func (l *LabelService) PutLabel(req domain.PutLabelReq) (*domain.PutLabelResp, error) {
	// todo: authorize req
	node, err := l.nodeRepo.Get(req.NodeId, req.Org)
	if err != nil {
		return nil, err
	}
	err = l.nodeRepo.PutLabel(*node, req.Label)
	if err != nil {
		return nil, err
	}
	return &domain.PutLabelResp{
		Node: *node,
	}, nil
}

func (l *LabelService) DeleteLabel(req domain.DeleteLabelReq) (*domain.DeleteLabelResp, error) {
	// todo: authorize req
	node, err := l.nodeRepo.Get(req.NodeId, req.Org)
	if err != nil {
		return nil, err
	}
	err = l.nodeRepo.DeleteLabel(*node, req.LabelKey)
	if err != nil {
		return nil, err
	}
	return &domain.DeleteLabelResp{
		Node: *node,
	}, nil
}
