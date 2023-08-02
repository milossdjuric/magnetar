package services

import (
	"github.com/c12s/magnetar/internal/domain"
)

type QueryService struct {
	nodeRepo domain.NodeRepo
}

func NewQueryService(nodeRepo domain.NodeRepo) (*QueryService, error) {
	return &QueryService{
		nodeRepo: nodeRepo,
	}, nil
}

func (q *QueryService) QueryNodes(req domain.QueryNodesReq) (*domain.QueryNodesResp, error) {
	nodeIds, err := q.nodeRepo.Query(req.Selector)
	if err != nil {
		return nil, err
	}
	nodes := make([]domain.Node, len(nodeIds))
	for i, nodeId := range nodeIds {
		node, err := q.nodeRepo.Get(nodeId)
		if err != nil {
			return nil, err
		}
		nodes[i] = *node
	}
	return &domain.QueryNodesResp{
		Nodes: nodes,
	}, nil
}
