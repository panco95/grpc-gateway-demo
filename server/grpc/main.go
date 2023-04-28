package main

import (
	"context"
	"google.golang.org/grpc"
	gen "grpc_gateway/gen/go/user"
	"log"
	"net"
)

// UserServerImpl will implement the service defined in protocol buffer definitions
type UserServerImpl struct {
	gen.UnimplementedUserServer
}

// Login is the implementation of RPC call defined in protocol definitions.
func (g *UserServerImpl) Login(ctx context.Context, request *gen.LoginRequest) (*gen.LoginReply, error) {
	return &gen.LoginReply{
		Token: request.Username + request.Password,
	}, nil
}
func main() {
	server := grpc.NewServer()
	gen.RegisterUserServer(server, &UserServerImpl{})
	if l, err := net.Listen("tcp", ":8080"); err != nil {
		log.Fatal("error in listening on port :8080", err)
	} else {
		// the gRPC server
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
