package domain

type Node struct {
	Id     NodeId
	Labels []Label
}

type NodeId struct {
	Value string
}

type QuerySelector []Query

type Query struct {
	LabelKey string
	ShouldBe ComparisonResult
	Value    string
}

type NodeRepo interface {
	Put(node Node) error
	Get(nodeId NodeId) (*Node, error)
	List() ([]Node, error)
	Query(selector QuerySelector) ([]NodeId, error)
	PutLabel(nodeId NodeId, label Label) error
	DeleteLabel(nodeId NodeId, labelKey string) error
}

type NodeMarshaller interface {
	Marshal(node Node) ([]byte, error)
	Unmarshal(nodeMarshalled []byte) (*Node, error)
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
	Selector QuerySelector
}

type QueryNodesResp struct {
	Nodes []Node
}
