package magnetar

import (
	"math"
	"strconv"
)

type RegistrationReq struct {
	Labels []Label
}

type RegistrationResp struct {
	NodeId string
}

type Label interface {
	Key() string
	Value() interface{}
	Compare(value string) ComparisonResult
}

func NewBoolLabel(key string, value bool) Label {
	return &boolLabel{
		key:   key,
		value: value,
	}
}

func NewFloat64Label(key string, value float64) Label {
	return &float64Label{
		key:   key,
		value: value,
	}
}

func NewStringLabel(key string, value string) Label {
	return &stringLabel{
		key:   key,
		value: value,
	}
}

type boolLabel struct {
	key   string
	value bool
}

func (b boolLabel) Key() string {
	return b.key
}

func (b boolLabel) Value() interface{} {
	return b.value
}

func (b boolLabel) Compare(value string) ComparisonResult {
	refValue, err := strconv.ParseBool(value)
	if err != nil {
		return CompResIncomparable
	}
	if b.value == refValue {
		return CompResEq
	}
	return CompResNeq
}

type float64Label struct {
	key   string
	value float64
}

func (f float64Label) Key() string {
	return f.key
}

func (f float64Label) Value() interface{} {
	return f.value
}

func (f float64Label) Compare(value string) ComparisonResult {
	refValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return CompResIncomparable
	}
	if math.Round(f.value*100)/100 == math.Round(refValue*100)/100 {
		return CompResEq
	}
	if f.value > refValue {
		return CompResGt
	}
	return CompResLt
}

type stringLabel struct {
	key   string
	value string
}

func (s stringLabel) Key() string {
	return s.key
}

func (s stringLabel) Value() interface{} {
	return s.value
}

func (s stringLabel) Compare(value string) ComparisonResult {
	if s.value == value {
		return CompResEq
	}
	return CompResNeq
}

type QuerySelector []Query

type Query struct {
	LabelKey string
	Value    string
	Expected []ComparisonResult
}

type ComparisonResult int8

const (
	CompResEq = iota
	CompResNeq
	CompResGt
	CompResLt
	CompResIncomparable
)
