package handlers

import (
	"github.com/c12s/magnetar/domain"
	"github.com/golang/protobuf/proto"
)

func MarshalResp(resp domain.RegistrationResp) ([]byte, error) {
	protoResp := &RegistrationResp{}
	protoResp = protoResp.fromDomain(resp)
	return proto.Marshal(protoResp)
}

func UnmarshalReq(req []byte) (*domain.RegistrationReq, error) {
	protoReq := RegistrationReq{}
	err := proto.Unmarshal(req, &protoReq)
	if err != nil {
		return nil, err
	}
	return protoReq.toDomain(), nil
}

func (x *RegistrationResp) fromDomain(resp domain.RegistrationResp) *RegistrationResp {
	return &RegistrationResp{
		NodeId: resp.NodeId,
	}
}

func (x *RegistrationReq) toDomain() *domain.RegistrationReq {
	return &domain.RegistrationReq{}
}
