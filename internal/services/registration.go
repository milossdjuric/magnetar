package services

import (
	"github.com/c12s/magnetar/internal/domain"
	"github.com/c12s/magnetar/pkg/magnetar"
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

func (r *RegistrationService) Register(req magnetar.RegistrationReq) (*magnetar.RegistrationResp, error) {
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
	res, err := r.nodeRepo.Query([]magnetar.Query{
		{
			LabelKey: "skey",
			Expected: []magnetar.ComparisonResult{magnetar.CompResEq},
			Value:    "abcd",
		},
		{
			LabelKey: "bkey",
			Expected: []magnetar.ComparisonResult{magnetar.CompResEq},
			Value:    "true",
		},
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(len(res))
	for _, id := range res {
		log.Println(id.Value)
	}

	return &magnetar.RegistrationResp{
		NodeId: node.Id.Value,
	}, nil
}

func generateNodeId() string {
	return uuid.NewString()
}
