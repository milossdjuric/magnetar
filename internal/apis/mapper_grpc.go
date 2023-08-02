package apis

import (
	"github.com/c12s/magnetar/internal/domain"
	"github.com/c12s/magnetar/pkg/api"
	"github.com/c12s/magnetar/pkg/magnetar"
)

func QueryNodesReq2Selector(req api.QueryNodesReq) (magnetar.QuerySelector, error) {
	res := make([]magnetar.Query, 0)
	for _, query := range req.Queries {
		shouldBe, err := magnetar.NewCompResultFromString(query.ShouldBe)
		if err != nil {
			return nil, err
		}
		resQuery := magnetar.Query{
			LabelKey: query.LabelKey,
			ShouldBe: shouldBe,
			Value:    query.Value,
		}
		res = append(res, resQuery)
	}
	return res, nil
}

func Nodes2QueryNodesResp(nodes []domain.Node) (*api.QueryNodesResp, error) {
	resp := &api.QueryNodesResp{
		Nodes: make([]*api.NodePretty, 0),
	}
	for _, node := range nodes {
		protoNode := &api.NodePretty{
			Id:     node.Id.Value,
			Labels: make([]*api.LabelPretty, 0),
		}
		for _, label := range node.Labels {
			protoLabel := &api.LabelPretty{
				Key:   label.Key(),
				Value: label.StringValue(),
			}
			protoNode.Labels = append(protoNode.Labels, protoLabel)
		}
		resp.Nodes = append(resp.Nodes, protoNode)
	}
	return resp, nil
}
