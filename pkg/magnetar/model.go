package magnetar

import (
	"errors"
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
	StringValue() string
	Compare(value string) (ComparisonResult, error)
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

func (b boolLabel) Compare(value string) (ComparisonResult, error) {
	refValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultCompRes, errors.New("incomparable")
	}
	if b.value == refValue {
		return CompResEq, nil
	}
	return CompResNeq, nil
}

func (b boolLabel) StringValue() string {
	return strconv.FormatBool(b.value)
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

func (f float64Label) Compare(value string) (ComparisonResult, error) {
	refValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultCompRes, errors.New("incomparable")
	}
	if math.Round(f.value*100)/100 == math.Round(refValue*100)/100 {
		return CompResEq, nil
	}
	if f.value > refValue {
		return CompResGt, nil
	}
	return CompResLt, nil
}

func (f float64Label) StringValue() string {
	return strconv.FormatFloat(f.value, 'f', 2, 64)
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

func (s stringLabel) StringValue() string {
	return s.value
}

func (s stringLabel) Compare(value string) (ComparisonResult, error) {
	if s.value == value {
		return CompResEq, nil
	}
	return CompResNeq, nil
}

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

type ComparisonResult int8

const (
	CompResEq = iota
	CompResNeq
	CompResGt
	CompResLt
)

func (c ComparisonResult) String() string {
	switch c {
	case CompResEq:
		return eqString
	case CompResNeq:
		return neqString
	case CompResGt:
		return gtString
	case CompResLt:
		return ltString
	default:
		return ""
	}
}

func NewCompResultFromString(value string) (ComparisonResult, error) {
	switch value {
	case eqString:
		return CompResEq, nil
	case neqString:
		return CompResNeq, nil
	case ltString:
		return CompResLt, nil
	case gtString:
		return CompResGt, nil
	default:
		return CompResEq, errors.New("invalid string")
	}
}

const (
	eqString  = "EQ"
	neqString = "NEQ"
	ltString  = "LT"
	gtString  = "GT"

	defaultCompRes = CompResEq
)
