package grpcserver

import (
	"context"
	"github.com/gnori-zon/go-tdd/scalingacceptancetests/domain/intersection"
)

type GreetServer struct {
	UnimplementedGreeterServer
}

func (g GreetServer) Greet(ctx context.Context, request *GreetRequest) (*GreetReply, error) {
	return &GreetReply{Message: intersection.Greet(request.Name)}, nil
}

func (g GreetServer) Curse(ctx context.Context, request *CurseRequest) (*CurseReply, error) {
	return &CurseReply{Message: intersection.Curse(request.Name)}, nil
}
