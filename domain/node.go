package domain

type Node struct {
	Id     NodeId
	Labels []Label
}

type NodeId struct {
	Value string
}

type Label interface {
	Key() string
	Value() interface{}
}

type label struct {
	key   string
	value interface{}
}

func NewBoolLabel(key string, value bool) Label {
	return &label{
		key:   key,
		value: value,
	}
}

func NewFloat64Label(key string, value float64) Label {
	return &label{
		key:   key,
		value: value,
	}
}

func NewStringLabel(key string, value string) Label {
	return &label{
		key:   key,
		value: value,
	}
}

func (b label) Key() string {
	return b.key
}

func (b label) Value() interface{} {
	return b.value
}

// todo: support operators other than == (<, >, !=)
type QuerySelector []Label

type NodeRepo interface {
	Put(node Node) error
	Get(nodeId NodeId) (*Node, error)
	Query(selector QuerySelector) ([]Node, error)
}
