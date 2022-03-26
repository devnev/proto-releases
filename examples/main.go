package main

import (
	"log"
	"net"

	_ "github.com/devnev/proto-releases/examples/init" // HACK, must be before proto imports
	examples "github.com/devnev/proto-releases/examples/server"
	alpha "github.com/devnev/proto-releases/examples/server/alpha"
	beta "github.com/devnev/proto-releases/examples/server/beta"
	stable "github.com/devnev/proto-releases/examples/server/stable"
	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
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

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err = srv.Serve(l); err != nil {
		log.Fatal(err)
	}
}
