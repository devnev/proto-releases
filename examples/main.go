// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			log.Printf("got %s: %s", info.FullMethod, protojson.Format(req.(proto.Message)))
			return handler(ctx, req)
		}),
		grpc.StreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			log.Printf("got %s", info.FullMethod)
			return handler(srv, ss)
		}),
	)

	var base baseServer
	examples.RegisterExampleServiceServer(srv, base)
	alpha.RegisterExampleServiceBaseServer(srv, base)
	beta.RegisterExampleServiceBaseServer(srv, base)
	stable.RegisterExampleServiceBaseServer(srv, base)
	reflection.Register(srv)

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

type baseServer struct {
	examples.UnsafeExampleServiceServer
}

func (baseServer) ExampleReleasedMethod(_ context.Context, e *examples.Example) (*examples.Example, error) {
	return &examples.Example{
		NotAnnotated:             12,
		Released:                 43,
		Previewed:                75,
		PreviewedReleasedRemoved: 99,
	}, nil
}

func (baseServer) ExampleUnreleaseMethod(_ context.Context, e *examples.Example) (*examples.Example, error) {
	return &examples.Example{
		NotAnnotated:             12,
		Released:                 43,
		Previewed:                75,
		PreviewedReleasedRemoved: 99,
	}, nil
}
