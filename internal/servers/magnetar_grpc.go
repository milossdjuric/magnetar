package servers

import (
	"context"
	"github.com/c12s/magnetar/internal/mappers/proto"
	"github.com/c12s/magnetar/internal/services"
	"github.com/c12s/magnetar/pkg/api"
)

type MagnetarGrpcServer struct {
	api.UnimplementedMagnetarServer
	nodeService  services.NodeService
	labelService services.LabelService
}

func NewMagnetarGrpcServer(nodeService services.NodeService, labelService services.LabelService) (api.MagnetarServer, error) {
	return &MagnetarGrpcServer{
		nodeService:  nodeService,
		labelService: labelService,
	}, nil
}

func (m *MagnetarGrpcServer) GetNode(ctx context.Context, req *api.GetNodeReq) (*api.GetNodeResp, error) {
	domainReq, err := proto.GetNodeReqToDomain(req)
	if err != nil {
		return nil, err
	}
	domainResp, err := m.nodeService.Get(*domainReq)
	if err != nil {
		return nil, err
	}
	return proto.GetNodeRespFromDomain(*domainResp)
}

func (m *MagnetarGrpcServer) ListNodes(ctx context.Context, req *api.ListNodesReq) (*api.ListNodesResp, error) {
	domainReq, err := proto.ListNodesReqToDomain(req)
	if err != nil {
		return nil, err
	}
	domainResp, err := m.nodeService.List(*domainReq)
	if err != nil {
		return nil, err
	}
	return proto.ListNodesRespFromDomain(*domainResp)
}

func (m *MagnetarGrpcServer) QueryNodes(ctx context.Context, req *api.QueryNodesReq) (*api.QueryNodesResp, error) {
	domainReq, err := proto.QueryNodesReqToDomain(req)
	if err != nil {
		return nil, err
	}
	domainResp, err := m.nodeService.Query(*domainReq)
	if err != nil {
		return nil, err
	}
	return proto.QueryNodesRespFromDomain(*domainResp)
}

func (m *MagnetarGrpcServer) PutBoolLabel(ctx context.Context, req *api.PutBoolLabelReq) (*api.PutLabelResp, error) {
	domainReq, err := proto.PutBoolLabelReqToDomain(req)
	if err != nil {
		return nil, err
	}
	domainResp, err := m.labelService.PutLabel(*domainReq)
	if err != nil {
		return nil, err
	}
	return proto.PutLabelRespFromDomain(*domainResp)
}

func (m *MagnetarGrpcServer) PutFloat64Label(ctx context.Context, req *api.PutFloat64LabelReq) (*api.PutLabelResp, error) {
	domainReq, err := proto.PutFloat64LabelReqToDomain(req)
	if err != nil {
		return nil, err
	}
	domainResp, err := m.labelService.PutLabel(*domainReq)
	if err != nil {
		return nil, err
	}
	return proto.PutLabelRespFromDomain(*domainResp)
}

func (m *MagnetarGrpcServer) PutStringLabel(ctx context.Context, req *api.PutStringLabelReq) (*api.PutLabelResp, error) {
	domainReq, err := proto.PutStringLabelReqToDomain(req)
	if err != nil {
		return nil, err
	}
	domainResp, err := m.labelService.PutLabel(*domainReq)
	if err != nil {
		return nil, err
	}
	return proto.PutLabelRespFromDomain(*domainResp)
}

func (m *MagnetarGrpcServer) DeleteLabel(ctx context.Context, req *api.DeleteLabelReq) (*api.DeleteLabelResp, error) {
	domainReq, err := proto.DeleteLabelReqToDomain(req)
	if err != nil {
		return nil, err
	}
	domainResp, err := m.labelService.DeleteLabel(*domainReq)
	if err != nil {
		return nil, err
	}
	return proto.DeleteLabelRespFromDomain(*domainResp)
}
