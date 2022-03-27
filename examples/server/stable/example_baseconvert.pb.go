// Code generated by protoc-gen-go-baseconvert. DO NOT EDIT.

package stable

import (
	context "context"
	server "github.com/devnev/proto-releases/examples/server"
)

func (m *Example) ToBase() *server.Example {
	msg := &server.Example{
		Released: m.GetReleased(),
	}
	return msg
}
func (m *Example) FromBase(b *server.Example) *Example {
	msg := &Example{
		Released: b.GetReleased(),
	}
	return msg
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
