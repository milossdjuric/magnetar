package storage

import (
	"context"
	mPb "github.com/c12s/scheme/magnetar"
)

type Storage interface {
	Reserve(ctx context.Context, req *mPb.ReserveMsg) (*mPb.ReserveRsp, error)
	List(ctx context.Context, req *mPb.ReserveMsg) (*mPb.ListRsp, error)
	Free(ctx context.Context, req *mPb.ReserveMsg) (*mPb.ReserveRsp, error)
	Count(ctx context.Context, id string) (int64, error)
	Store(ctx context.Context, id string, data, labels *mPb.DataMsg) error
	Query(ctx context.Context, req *mPb.DataMsg) (*mPb.ListRsp, error)
	Close()
}

type Metrics interface {
	Save(ctx context.Context, data *mPb.EventMsg, id, policy string) error
	Query(ctx context.Context, query []string, id string) error
	Init(ctx context.Context, id, policy string) error
	Close()
}
