package main

import (
	"context"
	"log"
	"net"
	// importing generated stubs
	"google.golang.org/grpc"
	gen "grpc_gateway/gen/go/user"
)

// UserServerImpl will implement the service defined in protocol buffer definitions
type UserServerImpl struct {
	gen.UnimplementedUserServer
}

// Login is the implementation of RPC call defined in protocol definitions.
// This will take HelloRequest message and return HelloReply
func (g *UserServerImpl) Login(ctx context.Context, request *gen.LoginRequest) (*gen.LoginReply, error) {
	return &gen.LoginReply{
		Token: request.Username + request.Password,
	}, nil
}
func main() {
	// create new gRPC server
	server := grpc.NewServer()
	// register the GreeterServerImpl on the gRPC server
	gen.RegisterUserServer(server, &UserServerImpl{})
	// start listening on port :8080 for a tcp connection
	if l, err := net.Listen("tcp", ":8080"); err != nil {
		log.Fatal("error in listening on port :8080", err)
	} else {
		// the gRPC server
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
