package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	golog "github.com/Vladroon22/GoLog"
	"github.com/Vladroon22/gRPC-service/internal/server"
	"github.com/Vladroon22/gRPC-service/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	logger := golog.New()
	s := grpc.NewServer()
	proto.RegisterServerServer(s, &server.Server{Logger: logger})
	logger.Infoln("Registration success")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Infoln("grpc-server is listening --> localhost:8080")

	if err := s.Serve(listener); err != nil {
		logger.Fatalln(err)
	}
	logger.Infoln("grpc-server is serving --> localhost:8080")

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit
}
