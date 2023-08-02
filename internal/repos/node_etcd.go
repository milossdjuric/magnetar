package repos

import (
	"context"
	"errors"
	"fmt"
	"github.com/c12s/magnetar/internal/domain"
	"github.com/c12s/magnetar/pkg/magnetar"
	"github.com/juliangruber/go-intersect"
	etcd "go.etcd.io/etcd/client/v3"
	"golang.org/x/exp/slices"
	"strings"
)

type nodeEtcdRepo struct {
	etcd       *etcd.Client
	marshaller magnetar.Marshaller
}

func NewNodeEtcdRepo(etcd *etcd.Client, marshaller magnetar.Marshaller) (domain.NodeRepo, error) {
	return &nodeEtcdRepo{
		etcd:       etcd,
		marshaller: marshaller,
	}, nil
}

func (n nodeEtcdRepo) Put(node domain.Node) error {
	err := n.putForGetting(node)
	if err != nil {
		return err
	}
	return n.putForQuerying(node)
}

func (n nodeEtcdRepo) putForGetting(node domain.Node) error {
	for _, label := range node.Labels {
		labelMarshalled, err := n.marshaller.MarshalLabel(label)
		if err != nil {
			return err
		}
		key := fmt.Sprintf("%s/labels/%s", node.Id.Value, label.Key())
		_, err = n.etcd.Put(context.TODO(), key, string(labelMarshalled))
		if err != nil {
			return err
		}
	}
	return nil
}

func (n nodeEtcdRepo) putForQuerying(node domain.Node) error {
	for _, label := range node.Labels {
		labelMarshalled, err := n.marshaller.MarshalLabel(label)
		if err != nil {
			return err
		}
		key := fmt.Sprintf("%s/%s", label.Key(), node.Id.Value)
		_, err = n.etcd.Put(context.TODO(), key, string(labelMarshalled))
		if err != nil {
			return err
		}
	}
	return nil
}

func (n nodeEtcdRepo) Get(nodeId domain.NodeId) (*domain.Node, error) {
	prefix := fmt.Sprintf("%s/labels/", nodeId.Value)
	resp, err := n.etcd.Get(context.TODO(), prefix, etcd.WithPrefix())
	if err != nil {
		return nil, err
	}
	if resp.Count == 0 {
		return nil, errors.New("node not found")
	}
	var labels []magnetar.Label
	for _, kv := range resp.Kvs {
		label, err := n.marshaller.UnmarshalLabel(kv.Value)
		if err != nil {
			return nil, err
		}
		labels = append(labels, label)
	}
	return &domain.Node{
		Id:     nodeId,
		Labels: labels,
	}, nil
}

func (n nodeEtcdRepo) Query(selector magnetar.QuerySelector) ([]domain.NodeId, error) {
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

func (n nodeEtcdRepo) query(query magnetar.Query) ([]domain.NodeId, error) {
	prefix := fmt.Sprintf("%s/", query.LabelKey)
	resp, err := n.etcd.Get(context.TODO(), prefix, etcd.WithPrefix())
	if err != nil {
		return nil, err
	}
	nodeIds := make([]domain.NodeId, 0)
	for _, kv := range resp.Kvs {
		nodeLabel, err := n.marshaller.UnmarshalLabel(kv.Value)
		if err != nil {
			return nil, err
		}
		compResult := nodeLabel.Compare(query.Value)
		if compResult == magnetar.CompResIncomparable {
			continue
		}
		if slices.Contains(query.Expected, compResult) {
			nodeId := strings.Split(string(kv.Key), "/")[1]
			nodeIds = append(nodeIds, domain.NodeId{
				Value: nodeId,
			})
		}
	}
	return nodeIds, nil
}
