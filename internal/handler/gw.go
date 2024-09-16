package handler

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	greetv1 "go.orx.me/echosrv/gen/greet/v1"
)

type GreetServer struct {
	greetv1.UnimplementedGreetServiceServer
}

func (s *GreetServer) Greet(context.Context, *greetv1.GreetRequest) (*greetv1.GreetResponse, error) {
	return &greetv1.GreetResponse{}, nil
}
func (s *GreetServer) Health(context.Context, *greetv1.HealthRequest) (*greetv1.HealthResponse, error) {
	return &greetv1.HealthResponse{}, nil
}

func MuxHandler() {
	m := runtime.NewServeMux()
	greetv1.RegisterGreetServiceHandlerServer(context.Background(), m, &GreetServer{})

}
