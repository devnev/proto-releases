// Code generated by protoc-gen-go-baseconvert. DO NOT EDIT.

package beta

import (
	context "context"
	server "github.com/devnev/proto-releases/examples/server"
)

func (m *Example) ToBase() *server.Example {
	msg := &server.Example{
		Released:  m.GetReleased(),
		Previewed: m.GetPreviewed(),
	}
	return msg
}
func (m *Example) FromBase(b *server.Example) *Example {
	if m != nil {
		m.Reset()
	} else {
		m = new(Example)
	}
	m.Released = b.GetReleased()
	m.Previewed = b.GetPreviewed()
	return m
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
