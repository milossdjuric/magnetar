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
	err := l.nodeRepo.PutLabel(req.NodeId, req.Label)
	if err != nil {
		return nil, err
	}
	node, err := l.nodeRepo.Get(req.NodeId)
	if err != nil {
		return nil, err
	}
	return &domain.PutLabelResp{
		Node: *node,
	}, nil
}
