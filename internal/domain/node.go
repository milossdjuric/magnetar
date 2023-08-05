package domain

import "github.com/c12s/magnetar/pkg/magnetar"

type NodeRepo interface {
	Put(node magnetar.Node) error
	Get(nodeId magnetar.NodeId) (*magnetar.Node, error)
	List() ([]magnetar.Node, error)
	Query(selector magnetar.QuerySelector) ([]magnetar.NodeId, error)
	PutLabel(nodeId magnetar.NodeId, label magnetar.Label) error
	DeleteLabel(nodeId magnetar.NodeId, labelKey string) error
}

type GetNodeReq struct {
	Id magnetar.NodeId
}

type GetNodeResp struct {
	Node magnetar.Node
}

type ListNodesReq struct {
}

type ListNodesResp struct {
	Nodes []magnetar.Node
}

type QueryNodesReq struct {
	Selector magnetar.QuerySelector
}

type QueryNodesResp struct {
	Nodes []magnetar.Node
}
