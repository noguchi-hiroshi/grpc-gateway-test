package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gateway "github.com/noguchi-hiroshi/grpc-gateway-test/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func run() error {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := "localhost:4949"
	err := gateway.RegisterUserServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":5000", mux)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
