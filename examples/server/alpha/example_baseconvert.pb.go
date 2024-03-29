// Code generated by protoc-gen-go-baseconvert. DO NOT EDIT.

package alpha

import (
	context "context"
	server "github.com/devnev/proto-releases/examples/server"
	grpc "google.golang.org/grpc"
)

func (m *Example) ToBase() *server.Example {
	msg := &server.Example{
		NotAnnotated:             m.GetNotAnnotated(),
		Released:                 m.GetReleased(),
		Previewed:                m.GetPreviewed(),
		PreviewedReleasedRemoved: m.GetPreviewedReleasedRemoved(),
	}
	return msg
}
func (m *Example) FromBase(b *server.Example) *Example {
	if m != nil {
		m.Reset()
	} else {
		m = new(Example)
	}
	m.NotAnnotated = b.GetNotAnnotated()
	m.Released = b.GetReleased()
	m.Previewed = b.GetPreviewed()
	m.PreviewedReleasedRemoved = b.GetPreviewedReleasedRemoved()
	return m
}

type baseExampleServiceServer struct {
	UnsafeExampleServiceServer
	base server.ExampleServiceServer
}

func RegisterExampleServiceBaseServer(s grpc.ServiceRegistrar, base server.ExampleServiceServer) {
	RegisterExampleServiceServer(s, baseExampleServiceServer{base: base})
}

func (s baseExampleServiceServer) ExampleUnreleaseMethod(ctx context.Context, in *Example) (*Example, error) {
	baseIn := in.ToBase()
	baseOut, err := s.base.ExampleUnreleaseMethod(ctx, baseIn)
	if err != nil {
		return nil, err
	}
	out := new(Example)
	out.FromBase(baseOut)
	return out, nil
}
func (s baseExampleServiceServer) ExampleReleasedMethod(ctx context.Context, in *Example) (*Example, error) {
	baseIn := in.ToBase()
	baseOut, err := s.base.ExampleReleasedMethod(ctx, baseIn)
	if err != nil {
		return nil, err
	}
	out := new(Example)
	out.FromBase(baseOut)
	return out, nil
}
