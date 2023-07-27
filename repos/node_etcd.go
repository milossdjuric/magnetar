package repos

import (
	"context"
	"errors"
	"fmt"
	"github.com/c12s/magnetar/domain"
	etcd "go.etcd.io/etcd/client/v3"
)

type nodeEtcdRepo struct {
	etcd *etcd.Client
}

func NewNodeEtcdRepo(etcd *etcd.Client) (domain.NodeRepo, error) {
	return &nodeEtcdRepo{
		etcd: etcd,
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
		labelMarshalled, err := MarshalLabel(label)
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
		labelMarshalled, err := MarshalLabel(label)
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
	var labels []domain.Label
	for _, kv := range resp.Kvs {
		label, err := UnmarshalLabel(kv.Value)
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

func (n nodeEtcdRepo) Query(selector domain.QuerySelector) ([]domain.Node, error) {
	return nil, nil
}
