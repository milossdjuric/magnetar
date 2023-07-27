package repos

import (
	"context"
	"errors"
	"fmt"
	"github.com/c12s/magnetar/internal/domain"
	"github.com/c12s/magnetar/pkg/magnetar"
	etcd "go.etcd.io/etcd/client/v3"
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

func (n nodeEtcdRepo) Query(selector domain.QuerySelector) ([]domain.Node, error) {
	return nil, nil
}
