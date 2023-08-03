package domain

import "github.com/c12s/magnetar/pkg/magnetar"

type Node struct {
	Id     NodeId
	Labels []magnetar.Label
}

type NodeId struct {
	Value string
}

type NodeRepo interface {
	Put(node Node) error
	Get(nodeId NodeId) (*Node, error)
	List() ([]Node, error)
	Query(selector magnetar.QuerySelector) ([]NodeId, error)
	PutLabel(nodeId NodeId, label magnetar.Label) error
	DeleteLabel(nodeId NodeId, labelKey string) error
}

type GetNodeReq struct {
	Id NodeId
}

type GetNodeResp struct {
	Node Node
}

type ListNodesReq struct {
}

type ListNodesResp struct {
	Nodes []Node
}

type QueryNodesReq struct {
	Selector magnetar.QuerySelector
}

type QueryNodesResp struct {
	Nodes []Node
}
