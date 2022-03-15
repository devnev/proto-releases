// Code generated by protoc-gen-baseconvert. DO NOT EDIT.

package stable2

import (
	context "context"
	server "github.com/devnev/proto-releases/examples/server"
)

func (m *Example) ToBase() *server.Example {
	msg := &server.Example{
		Released:                 m.GetReleased(),
		PreviewedReleasedRemoved: m.GetPreviewedReleasedRemoved(),
	}
	return msg
}
func (m *Example) FromBase(b *server.Example) {
	m.Reset()
	m.Released = b.GetReleased()
	m.PreviewedReleasedRemoved = b.GetPreviewedReleasedRemoved()
}

type BaseExampleServiceServer struct {
	UnsafeExampleServiceServer
	Base server.ExampleServiceServer
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
