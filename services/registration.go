package services

import (
	"github.com/c12s/magnetar/domain"
	"github.com/google/uuid"
	"log"
)

type RegistrationService struct {
}

func NewRegistrationService() *RegistrationService {
	return &RegistrationService{}
}

func (rs *RegistrationService) Register(req domain.RegistrationReq) (*domain.RegistrationResp, error) {
	nodeId := uuid.NewString()
	log.Println(nodeId)
	// todo: save node id and labels to db
	return &domain.RegistrationResp{
		NodeId: nodeId,
	}, nil
}
