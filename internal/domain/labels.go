package domain

import "github.com/c12s/magnetar/pkg/magnetar"

type PutLabelReq struct {
	NodeId NodeId
	Label  magnetar.Label
}

type PutLabelResp struct {
	Node Node
}
