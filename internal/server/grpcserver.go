package server

import (
	"context"

	golog "github.com/Vladroon22/GoLog"
	"github.com/Vladroon22/gRPC-service/pkg/proto"
)

type Server struct {
	proto.UnimplementedServerServer
	Logger *golog.Logger
}

func (s *Server) BroadCast(ctx context.Context, client *proto.ClientText) (*proto.ServerText, error) {
	return &proto.ServerText{Response: client.Text}, nil
}
