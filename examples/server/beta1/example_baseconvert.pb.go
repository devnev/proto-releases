// Code generated by protoc-gen-baseconvert. DO NOT EDIT.

package beta1

import (
	server "github.com/devnev/proto-releases/examples/server"
)

func (m *Example) ToBase() *server.Example {
	msg := &server.Example{
		PreviewedReleasedRemoved: m.GetPreviewedReleasedRemoved(),
	}
	return msg
}
func (m *Example) FromBase(b *server.Example) {
	m.Reset()
	m.PreviewedReleasedRemoved = b.GetPreviewedReleasedRemoved()
}
