// Code generated by protoc-gen-go-baseconvert. DO NOT EDIT.

package alpha

import (
	context "context"
	server "github.com/devnev/proto-releases/examples/server"
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
	msg := &Example{
		NotAnnotated:             b.GetNotAnnotated(),
		Released:                 b.GetReleased(),
		Previewed:                b.GetPreviewed(),
		PreviewedReleasedRemoved: b.GetPreviewedReleasedRemoved(),
	}
	return msg
}

type BaseExampleServiceServer struct {
	UnsafeExampleServiceServer
	Base server.ExampleServiceServer
}

func (s BaseExampleServiceServer) ExampleUnreleaseMethod(ctx context.Context, in *Example) (*Example, error) {
	baseIn := in.ToBase()
	baseOut, err := s.Base.ExampleUnreleaseMethod(ctx, baseIn)
	if err != nil {
		return nil, err
	}
	out := new(Example)
	out.FromBase(baseOut)
	return out, nil
}
func (s BaseExampleServiceServer) ExampleReleasedMethod(ctx context.Context, in *Example) (*Example, error) {
	baseIn := in.ToBase()
	baseOut, err := s.Base.ExampleReleasedMethod(ctx, baseIn)
	if err != nil {
		return nil, err
	}
	out := new(Example)
	out.FromBase(baseOut)
	return out, nil
}
