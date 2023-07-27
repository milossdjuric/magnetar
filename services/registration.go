package services

import (
	"github.com/c12s/magnetar/domain"
	"github.com/google/uuid"
	"log"
)

type RegistrationService struct {
	nodeRepo domain.NodeRepo
}

func NewRegistrationService(nodeRepo domain.NodeRepo) (*RegistrationService, error) {
	return &RegistrationService{
		nodeRepo: nodeRepo,
	}, nil
}

func (r *RegistrationService) Register(req domain.RegistrationReq) (*domain.RegistrationResp, error) {
	node := domain.Node{
		Id: domain.NodeId{
			Value: generateNodeId(),
		},
		Labels: req.Labels,
	}

	err := r.nodeRepo.Put(node)
	if err != nil {
		return nil, err
	}

	// todo: delete this later
	log.Println("TEST GETTING THE NODE")
	fetchedNode, err := r.nodeRepo.Get(node.Id)
	log.Println(err)
	log.Println(fetchedNode)

	return &domain.RegistrationResp{
		NodeId: node.Id.Value,
	}, nil
}

func generateNodeId() string {
	return uuid.NewString()
}
