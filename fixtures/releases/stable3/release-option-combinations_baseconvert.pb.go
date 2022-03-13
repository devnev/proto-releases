// Code generated by protoc-gen-baseconvert. DO NOT EDIT.

package fixtures

import (
	context "context"
	fixtures "github.com/devnev/proto-releases/fixtures"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (m *EmptyMessageReleased) ToBase() *fixtures.EmptyMessageReleased {
	msg := &fixtures.EmptyMessageReleased{}
	return msg
}
func (m *EmptyMessageReleased) FromBase(b *fixtures.EmptyMessageReleased) {
	m.Reset()
}
func (m *EmptyMessagePreviewedThenReleased) ToBase() *fixtures.EmptyMessagePreviewedThenReleased {
	msg := &fixtures.EmptyMessagePreviewedThenReleased{}
	return msg
}
func (m *EmptyMessagePreviewedThenReleased) FromBase(b *fixtures.EmptyMessagePreviewedThenReleased) {
	m.Reset()
}
func (m *MessageWithReleasedField) ToBase() *fixtures.MessageWithReleasedField {
	msg := &fixtures.MessageWithReleasedField{
		Released: m.GetReleased(),
	}
	return msg
}
func (m *MessageWithReleasedField) FromBase(b *fixtures.MessageWithReleasedField) {
	m.Reset()
	m.Released = b.GetReleased()
}
func (m *MessageWithReleasedOneofItem) ToBase() *fixtures.MessageWithReleasedOneofItem {
	msg := &fixtures.MessageWithReleasedOneofItem{}
	switch o := m.GetOneofWithItem().(type) {
	case *MessageWithReleasedOneofItem_ReleasedOneofItem:
		msg.OneofWithItem = (*fixtures.MessageWithReleasedOneofItem_ReleasedOneofItem)(o)
	}
	return msg
}
func (m *MessageWithReleasedOneofItem) FromBase(b *fixtures.MessageWithReleasedOneofItem) {
	m.Reset()
	switch o := b.GetOneofWithItem().(type) {
	case *fixtures.MessageWithReleasedOneofItem_ReleasedOneofItem:
		m.OneofWithItem = (*MessageWithReleasedOneofItem_ReleasedOneofItem)(o)
	}
}
func (m *MessageNotAnnotated) ToBase() *fixtures.MessageNotAnnotated {
	msg := &fixtures.MessageNotAnnotated{
		Released:              m.GetReleased(),
		PreviewedThenReleased: m.GetPreviewedThenReleased(),
	}
	switch o := m.GetNotAnnotatedOneof().(type) {
	case *MessageNotAnnotated_OneofItemNotAnnotated:
		msg.NotAnnotatedOneof = (*fixtures.MessageNotAnnotated_OneofItemNotAnnotated)(o)
	}
	return msg
}
func (m *MessageNotAnnotated) FromBase(b *fixtures.MessageNotAnnotated) {
	m.Reset()
	m.Released = b.GetReleased()
	m.PreviewedThenReleased = b.GetPreviewedThenReleased()
	switch o := b.GetNotAnnotatedOneof().(type) {
	case *fixtures.MessageNotAnnotated_OneofItemNotAnnotated:
		m.NotAnnotatedOneof = (*MessageNotAnnotated_OneofItemNotAnnotated)(o)
	}
}

type BaseTestServiceServer struct {
	UnsafeTestServiceServer
	Base fixtures.TestServiceServer
}

func (s BaseTestServiceServer) EmptyMethodReleased(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	return s.Base.EmptyMethodReleased(ctx, in)
}
func (s BaseTestServiceServer) MethodReleased(ctx context.Context, in *MessageNotAnnotated) (*MessageNotAnnotated, error) {
	baseIn := in.ToBase()
	baseOut, err := s.Base.MethodReleased(ctx, baseIn)
	if err != nil {
		return nil, err
	}
	out := new(MessageNotAnnotated)
	out.FromBase(baseOut)
	return out, nil
}
