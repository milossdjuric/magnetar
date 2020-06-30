package service

import (
	"context"
	"fmt"
	"github.com/c12s/magnetar/storage"
	"github.com/c12s/magnetar/sync"
	mPb "github.com/c12s/scheme/magnetar"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	db   storage.Storage
	mdb  storage.Metrics
	sync syncer.Syncer
}

func (s *Server) Reserve(ctx context.Context, req *mPb.ReserveMsg) (*mPb.ReserveRsp, error) {
	return s.db.Reserve(ctx, req)
}

func (s *Server) Free(ctx context.Context, req *mPb.ReserveMsg) (*mPb.ReserveRsp, error) {
	return s.db.Free(ctx, req)
}

func (s *Server) List(ctx context.Context, req *mPb.ReserveMsg) (*mPb.ListRsp, error) {
	return s.db.List(ctx, req)
}

func (s *Server) Query(ctx context.Context, req *mPb.DataMsg) (*mPb.ListRsp, error) {
	return s.db.Query(ctx, req)
}

func (s *Server) Health(ctx context.Context) {
	s.sync.Sub(func(msg *mPb.EventMsg) {
		general := msg.Data["general"].Data
		fmt.Println("KEY ARRIVED: ", general["key"])
		if isUsed(general["key"]) {
			fmt.Println("KEY USED: ", allocatedKey(general["key"]))

			//store host data
			err := s.db.Store(ctx, allocatedKey(general["key"]), msg.Data["host"], nil)
			if err != nil {
				return
			}

			//store metrics data
			err = s.mdb.Save(ctx, msg, general["key"], general["policy"])
			if err != nil {
				return
			}
		} else {
			// if node is not used then check if it is reserved
			reserveKey := reservedKey(general["key"])
			resp, err := s.db.Count(ctx, reserveKey)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			key := nodeKey(general["key"])
			// if it is not reserved, than store it into db with lease
			if resp == 0 {
				// store host data
				err = s.db.Store(ctx, key, msg.Data["host"], msg.Data["labels"])
				if err != nil {
					return
				}

				// store metriccs data
				err = s.mdb.Save(ctx, msg, general["key"], general["policy"])
				if err != nil {
					return
				}

				fmt.Println("KEY NOT RESERVED: ", key)
				return
			}
			fmt.Println("KEY RESERVED: ", general["key"])
		}
	})
}

func Run(s storage.Storage, m storage.Metrics, sync syncer.Syncer, address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	magnetarServer := &Server{
		db:   s,
		mdb:  m,
		sync: sync,
	}

	ctx, cancel := context.WithCancel(context.Background())
	magnetarServer.Health(ctx)
	defer s.Close()
	defer m.Close()
	defer cancel()

	fmt.Println("MagnetarService RPC Started")
	mPb.RegisterMagnetarServiceServer(server, magnetarServer)
	server.Serve(lis)
}
