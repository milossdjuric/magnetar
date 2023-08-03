package handlers

import (
	"context"
	"github.com/c12s/magnetar/internal/services"
	"github.com/c12s/magnetar/pkg/proto"
)

type MagnetarGrpcServer struct {
	proto.UnimplementedMagnetarServer
	nodeService  services.NodeService
	labelService services.LabelService
}

func NewMagnetarGrpcServer(nodeService services.NodeService, labelService services.LabelService) (proto.MagnetarServer, error) {
	return &MagnetarGrpcServer{
		nodeService:  nodeService,
		labelService: labelService,
	}, nil
}

func (m *MagnetarGrpcServer) GetNode(ctx context.Context, req *proto.GetNodeReq) (*proto.GetNodeResp, error) {
	domainReq, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	domainResp, err := m.nodeService.Get(*domainReq)
	if err != nil {
		return nil, err
	}
	resp := &proto.GetNodeResp{}
	return resp.FromDomain(*domainResp)
}

func (m *MagnetarGrpcServer) ListNodes(ctx context.Context, req *proto.ListNodesReq) (*proto.ListNodesResp, error) {
	domainReq, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	domainResp, err := m.nodeService.List(*domainReq)
	if err != nil {
		return nil, err
	}
	resp := &proto.ListNodesResp{}
	return resp.FromDomain(*domainResp)
}

func (m *MagnetarGrpcServer) QueryNodes(ctx context.Context, req *proto.QueryNodesReq) (*proto.QueryNodesResp, error) {
	domainReq, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	domainResp, err := m.nodeService.Query(*domainReq)
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

func (m *MagnetarGrpcServer) DeleteLabel(ctx context.Context, req *proto.DeleteLabelReq) (*proto.DeleteLabelResp, error) {
	domainReq, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	domainResp, err := m.labelService.DeleteLabel(*domainReq)
	if err != nil {
		return nil, err
	}
	resp := proto.DeleteLabelResp{}
	return resp.FromDomain(*domainResp)
}
