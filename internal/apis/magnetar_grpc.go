package apis

import (
	"context"
	"github.com/c12s/magnetar/internal/services"
	"github.com/c12s/magnetar/pkg/proto"
)

type MagnetarGrpcServer struct {
	proto.UnimplementedMagnetarServer
	queryService services.QueryService
	labelService services.LabelService
}

func NewMagnetarGrpcServer(queryService services.QueryService, labelService services.LabelService) (proto.MagnetarServer, error) {
	return &MagnetarGrpcServer{
		queryService: queryService,
		labelService: labelService,
	}, nil
}

func (m *MagnetarGrpcServer) QueryNodes(ctx context.Context, req *proto.QueryNodesReq) (*proto.QueryNodesResp, error) {
	domainReq, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	domainResp, err := m.queryService.QueryNodes(*domainReq)
	if err != nil {
		return nil, err
	}
	resp := &proto.QueryNodesResp{}
	return resp.FromDomain(*domainResp)
}

func (m *MagnetarGrpcServer) PutLabel(ctx context.Context, req *proto.PutLabelReq) (*proto.PutLabelResp, error) {
	domainReq, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	domainResp, err := m.labelService.PutLabel(*domainReq)
	if err != nil {
		return nil, err
	}
	resp := proto.PutLabelResp{}
	return resp.FromDomain(*domainResp)
}
