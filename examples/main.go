package main

import (
	"context"
	"log"
	"net"

	examples "github.com/devnev/proto-releases/examples/server"
	alpha "github.com/devnev/proto-releases/examples/server/alpha"
	beta "github.com/devnev/proto-releases/examples/server/beta"
	stable "github.com/devnev/proto-releases/examples/server/stable"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			log.Printf("got %s", info.FullMethod)
			return handler(ctx, req)
		}),
		grpc.StreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			log.Printf("got %s", info.FullMethod)
			return handler(srv, ss)
		}),
	)

	reflection.Register(srv)
	var base examples.ExampleServiceServer = examples.UnimplementedExampleServiceServer{}
	examples.RegisterExampleServiceServer(srv, base)
	alpha.RegisterExampleServiceServer(srv, alpha.BaseExampleServiceServer{
		Base: base,
	})
	beta.RegisterExampleServiceServer(srv, beta.BaseExampleServiceServer{
		Base: base,
	})
	stable.RegisterExampleServiceServer(srv, stable.BaseExampleServiceServer{
		Base: base,
	})

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Listening on ", l.Addr())
	log.Printf("Try running: grpcurl -plaintext -v %s example.alpha.ExampleService/ExampleUnreleaseMethod", l.Addr())
	if err = srv.Serve(l); err != nil {
		log.Fatal(err)
	}
}
