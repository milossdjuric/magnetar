package startup

import (
	"github.com/c12s/magnetar/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func startServer(address string, magnetarServer api.MagnetarServer) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	api.RegisterMagnetarServer(s, magnetarServer)
	reflection.Register(s)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return s, nil
}
