package pkg

import (
	"github.com/c12s/magnetar/internal/domain"
	"github.com/c12s/magnetar/pkg/magnetar"
)

type Marshaller interface {
	MarshalRegistrationReq(req magnetar.RegistrationReq) ([]byte, error)
	UnmarshalRegistrationReq(reqMarshalled []byte) (*magnetar.RegistrationReq, error)
	MarshalRegistrationResp(resp magnetar.RegistrationResp) ([]byte, error)
	UnmarshalRegistrationResp(resp []byte) (*magnetar.RegistrationResp, error)
	MarshalLabel(label magnetar.Label) ([]byte, error)
	UnmarshalLabel(labelMarshalled []byte) (magnetar.Label, error)
	MarshalNode(node domain.Node) ([]byte, error)
	UnmarshalNode(nodeMarshalled []byte) (*domain.Node, error)
}
