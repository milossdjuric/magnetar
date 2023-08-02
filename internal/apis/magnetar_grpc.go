package apis

import (
	"context"
	"github.com/c12s/magnetar/internal/services"
	"github.com/c12s/magnetar/pkg/api"
)

type MagnetarGrpcServer struct {
	api.UnimplementedMagnetarServer
	service services.QueryService
}

func NewMagnetarGrpcServer(service services.QueryService) (api.MagnetarServer, error) {
	return &MagnetarGrpcServer{
		service: service,
	}, nil
}

func (m *MagnetarGrpcServer) QueryNodes(ctx context.Context, req *api.QueryNodesReq) (*api.QueryNodesResp, error) {
	selector, err := QueryNodesReq2Selector(req)
	if err != nil {
		return nil, err
	}
	nodes, err := m.service.QueryNodes(selector)
	if err != nil {
		return nil, err
	}
	return Nodes2QueryNodesResp(nodes)
}
