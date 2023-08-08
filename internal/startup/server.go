package startup

import (
	"github.com/c12s/magnetar/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func startServer(address string, server proto.MagnetarServer) chan bool {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterMagnetarServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	reflection.Register(s)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	var serverStoppedCh chan bool
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGINT)
		<-quit

		s.GracefulStop()
		serverStoppedCh <- true
	}()

	return serverStoppedCh
}
