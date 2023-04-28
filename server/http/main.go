package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	gen "grpc_gateway/gen/go/user"
	"log"
)

func main() {
	mux := runtime.NewServeMux()
	err := gen.RegisterUserHandlerFromEndpoint(context.Background(), mux, "localhost:8080", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Group("v1/*name").Any("", gin.WrapH(mux))
	err = r.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
