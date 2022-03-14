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
func (m *EmptyMessagePreviewed) ToBase() *fixtures.EmptyMessagePreviewed {
	msg := &fixtures.EmptyMessagePreviewed{}
	return msg
}
func (m *EmptyMessagePreviewed) FromBase(b *fixtures.EmptyMessagePreviewed) {
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
func (m *MessageWithPreviewField) ToBase() *fixtures.MessageWithPreviewField {
	msg := &fixtures.MessageWithPreviewField{
		Released: m.GetReleased(),
	}
	return msg
}
func (m *MessageWithPreviewField) FromBase(b *fixtures.MessageWithPreviewField) {
	m.Reset()
	m.Released = b.GetReleased()
}
func (m *MessageWithReleasedOneofItem) ToBase() *fixtures.MessageWithReleasedOneofItem {
	msg := &fixtures.MessageWithReleasedOneofItem{}
	switch o := m.GetOneofWithItem().(type) {
	case *MessageWithReleasedOneofItem_ReleasedOneofItem:
		msg.OneofWithItem = o.ToBase()
	}
	return msg
}
func (m *MessageWithReleasedOneofItem) FromBase(b *fixtures.MessageWithReleasedOneofItem) {
	m.Reset()
	switch o := b.GetOneofWithItem().(type) {
	case *fixtures.MessageWithReleasedOneofItem_ReleasedOneofItem:
		msg := new(MessageWithReleasedOneofItem_ReleasedOneofItem)
		msg.FromBase(o)
		m.OneofWithItem = msg
	}
}
func (m *MessageWithReleasedOneofItem_ReleasedOneofItem) ToBase() *fixtures.MessageWithReleasedOneofItem_ReleasedOneofItem {
	return (*fixtures.MessageWithReleasedOneofItem_ReleasedOneofItem)(m)
}
func (m *MessageWithReleasedOneofItem_ReleasedOneofItem) FromBase(b *fixtures.MessageWithReleasedOneofItem_ReleasedOneofItem) {
	*m = *(*MessageWithReleasedOneofItem_ReleasedOneofItem)(b)
}
func (m *MessageWithOneofWithMessages) ToBase() *fixtures.MessageWithOneofWithMessages {
	msg := &fixtures.MessageWithOneofWithMessages{}
	switch o := m.GetOneofWithMessage().(type) {
	case *MessageWithOneofWithMessages_MessageWithReleaseAnnotation:
		msg.OneofWithMessage = o.ToBase()
	}
	return msg
}
func (m *MessageWithOneofWithMessages) FromBase(b *fixtures.MessageWithOneofWithMessages) {
	m.Reset()
	switch o := b.GetOneofWithMessage().(type) {
	case *fixtures.MessageWithOneofWithMessages_MessageWithReleaseAnnotation:
		msg := new(MessageWithOneofWithMessages_MessageWithReleaseAnnotation)
		msg.FromBase(o)
		m.OneofWithMessage = msg
	}
}
func (m *MessageWithOneofWithMessages_MessageWithReleaseAnnotation) ToBase() *fixtures.MessageWithOneofWithMessages_MessageWithReleaseAnnotation {
	return &fixtures.MessageWithOneofWithMessages_MessageWithReleaseAnnotation{
		MessageWithReleaseAnnotation: m.MessageWithReleaseAnnotation.ToBase(),
	}
}
func (m *MessageWithOneofWithMessages_MessageWithReleaseAnnotation) FromBase(b *fixtures.MessageWithOneofWithMessages_MessageWithReleaseAnnotation) {
	m.MessageWithReleaseAnnotation = new(MessageWithReleasedField)
	m.MessageWithReleaseAnnotation.FromBase(b.MessageWithReleaseAnnotation)
}
func (m *MessageNotAnnotated) ToBase() *fixtures.MessageNotAnnotated {
	msg := &fixtures.MessageNotAnnotated{
		Released:              m.GetReleased(),
		Previewed:             m.GetPreviewed(),
		PreviewedThenReleased: m.GetPreviewedThenReleased(),
	}
	switch o := m.GetNotAnnotatedOneof().(type) {
	case *MessageNotAnnotated_OneofItemNotAnnotated:
		msg.NotAnnotatedOneof = o.ToBase()
	}
	return msg
}
func (m *MessageNotAnnotated) FromBase(b *fixtures.MessageNotAnnotated) {
	m.Reset()
	m.Released = b.GetReleased()
	m.Previewed = b.GetPreviewed()
	m.PreviewedThenReleased = b.GetPreviewedThenReleased()
	switch o := b.GetNotAnnotatedOneof().(type) {
	case *fixtures.MessageNotAnnotated_OneofItemNotAnnotated:
		msg := new(MessageNotAnnotated_OneofItemNotAnnotated)
		msg.FromBase(o)
		m.NotAnnotatedOneof = msg
	}
}
func (m *MessageNotAnnotated_OneofItemNotAnnotated) ToBase() *fixtures.MessageNotAnnotated_OneofItemNotAnnotated {
	return (*fixtures.MessageNotAnnotated_OneofItemNotAnnotated)(m)
}
func (m *MessageNotAnnotated_OneofItemNotAnnotated) FromBase(b *fixtures.MessageNotAnnotated_OneofItemNotAnnotated) {
	*m = *(*MessageNotAnnotated_OneofItemNotAnnotated)(b)
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
