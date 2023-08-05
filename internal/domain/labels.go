package domain

import "github.com/c12s/magnetar/pkg/magnetar"

type PutLabelReq struct {
	NodeId magnetar.NodeId
	Label  magnetar.Label
}

type PutLabelResp struct {
	Node magnetar.Node
}

type DeleteLabelReq struct {
	NodeId   magnetar.NodeId
	LabelKey string
}

type DeleteLabelResp struct {
	Node magnetar.Node
}
