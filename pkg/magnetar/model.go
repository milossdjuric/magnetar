package magnetar

type RegistrationReq struct {
	Labels []Label
}

type RegistrationResp struct {
	NodeId string
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
