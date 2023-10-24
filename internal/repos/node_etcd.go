package repos

import (
	"context"
	"errors"
	"fmt"
	"github.com/c12s/magnetar/internal/domain"
	"github.com/juliangruber/go-intersect"
	etcd "go.etcd.io/etcd/client/v3"
	"golang.org/x/exp/slices"
	"log"
	"strings"
)

type nodeEtcdRepo struct {
	etcd            *etcd.Client
	nodeMarshaller  domain.NodeMarshaller
	labelMarshaller domain.LabelMarshaller
}

func NewNodeEtcdRepo(etcd *etcd.Client, nodeMarshaller domain.NodeMarshaller, labelMarshaller domain.LabelMarshaller) (domain.NodeRepo, error) {
	return &nodeEtcdRepo{
		etcd:            etcd,
		nodeMarshaller:  nodeMarshaller,
		labelMarshaller: labelMarshaller,
	}, nil
}

func (n nodeEtcdRepo) Put(node domain.Node) error {
	err := n.putNodeForGetting(node)
	if err != nil {
		return err
	}
	return n.putNodeForQuerying(node)
}

func (n nodeEtcdRepo) Get(nodeId domain.NodeId) (*domain.Node, error) {
	key := fmt.Sprintf("nodes/%s", nodeId.Value)
	resp, err := n.etcd.Get(context.TODO(), key)
	if err != nil {
		return nil, err
	}
	if resp.Count == 0 {
		return nil, errors.New("node not found")
	}
	return n.nodeMarshaller.Unmarshal(resp.Kvs[0].Value)
}

func (n nodeEtcdRepo) List() ([]domain.Node, error) {
	nodes := make([]domain.Node, 0)
	prefix := "nodes/"
	resp, err := n.etcd.Get(context.TODO(), prefix, etcd.WithPrefix())
	if err != nil {
		return nil, err
	}
	for _, kv := range resp.Kvs {
		node, err := n.nodeMarshaller.Unmarshal(kv.Value)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, *node)
	}
	return nodes, nil
}

func (n nodeEtcdRepo) Query(selector domain.QuerySelector) ([]domain.NodeId, error) {
	if len(selector) == 0 {
		return nil, errors.New("empty selector")
	}
	nodes := make([]domain.NodeId, 0)
	for i, query := range selector {
		currNodes, err := n.query(query)
		if err != nil {
			return nil, err
		}
		if i == 0 {
			nodes = currNodes
		} else {
			intersection := intersect.Simple(nodes, currNodes)
			nodes = make([]domain.NodeId, len(intersection))
			for i, node := range intersection {
				nodes[i] = node.(domain.NodeId)
			}
		}
	}
	return nodes, nil
}

func (n nodeEtcdRepo) PutLabel(nodeId domain.NodeId, label domain.Label) error {
	err := n.putLabelForGetting(nodeId, label)
	if err != nil {
		return err
	}
	return n.putLabelForQuerying(nodeId, label)
}

func (n nodeEtcdRepo) DeleteLabel(nodeId domain.NodeId, labelKey string) error {
	err := n.deleteLabelForGetting(nodeId, labelKey)
	if err != nil {
		return err
	}
	return n.deleteLabelForQuerying(nodeId, labelKey)
}

func (n nodeEtcdRepo) putNodeForGetting(node domain.Node) error {
	nodeMarshalled, err := n.nodeMarshaller.Marshal(node)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("nodes/%s", node.Id.Value)
	_, err = n.etcd.Put(context.TODO(), key, string(nodeMarshalled))
	return err
}

func (n nodeEtcdRepo) putLabelForGetting(nodeId domain.NodeId, label domain.Label) error {
	node, err := n.Get(nodeId)
	if err != nil {
		return err
	}
	labelIndex := -1
	for i, nodeLabel := range node.Labels {
		if nodeLabel.Key() == label.Key() {
			labelIndex = i
		}
	}
	if labelIndex >= 0 {
		node.Labels[labelIndex] = label
	} else {
		node.Labels = append(node.Labels, label)
	}
	return n.putNodeForGetting(*node)
}

func (n nodeEtcdRepo) deleteLabelForGetting(nodeId domain.NodeId, labelKey string) error {
	node, err := n.Get(nodeId)
	if err != nil {
		return err
	}
	labelIndex := -1
	for i, nodeLabel := range node.Labels {
		if nodeLabel.Key() == labelKey {
			labelIndex = i
		}
	}
	if labelIndex >= 0 {
		node.Labels = slices.Delete(node.Labels, labelIndex, labelIndex+1)
		return n.putNodeForGetting(*node)
	}
	return nil
}

func (n nodeEtcdRepo) putNodeForQuerying(node domain.Node) error {
	for _, label := range node.Labels {
		err := n.putLabelForQuerying(node.Id, label)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n nodeEtcdRepo) putLabelForQuerying(nodeId domain.NodeId, label domain.Label) error {
	labelMarshalled, err := n.labelMarshaller.Marshal(label)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("%s/%s", label.Key(), nodeId.Value)
	_, err = n.etcd.Put(context.TODO(), key, string(labelMarshalled))
	return err
}

func (n nodeEtcdRepo) deleteLabelForQuerying(nodeId domain.NodeId, labelKey string) error {
	key := fmt.Sprintf("%s/%s", labelKey, nodeId.Value)
	_, err := n.etcd.Delete(context.TODO(), key)
	return err
}

func (n nodeEtcdRepo) query(query domain.Query) ([]domain.NodeId, error) {
	prefix := fmt.Sprintf("%s/", query.LabelKey)
	resp, err := n.etcd.Get(context.TODO(), prefix, etcd.WithPrefix())
	if err != nil {
		return nil, err
	}
	nodeIds := make([]domain.NodeId, 0)
	for _, kv := range resp.Kvs {
		nodeLabel, err := n.labelMarshaller.Unmarshal(kv.Value)
		if err != nil {
			return nil, err
		}
		cmpResult, err := nodeLabel.Compare(query.Value)
		if err != nil {
			log.Println(err)
			continue
		}
		if slices.Contains(cmpResult, query.ShouldBe) {
			nodeId := strings.Split(string(kv.Key), "/")[1]
			nodeIds = append(nodeIds, domain.NodeId{
				Value: nodeId,
			})
		}
	}
	return nodeIds, nil
}
