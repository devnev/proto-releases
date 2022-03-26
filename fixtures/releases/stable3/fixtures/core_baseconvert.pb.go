// Code generated by protoc-gen-go-baseconvert. DO NOT EDIT.

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
func (m *EmptyMessageReleased) FromBase(b *fixtures.EmptyMessageReleased) *EmptyMessageReleased {
	msg := &EmptyMessageReleased{}
	return msg
}
func (m *EmptyMessagePreviewedThenReleased) ToBase() *fixtures.EmptyMessagePreviewedThenReleased {
	msg := &fixtures.EmptyMessagePreviewedThenReleased{}
	return msg
}
func (m *EmptyMessagePreviewedThenReleased) FromBase(b *fixtures.EmptyMessagePreviewedThenReleased) *EmptyMessagePreviewedThenReleased {
	msg := &EmptyMessagePreviewedThenReleased{}
	return msg
}
func (m *MessageWithReleasedField) ToBase() *fixtures.MessageWithReleasedField {
	msg := &fixtures.MessageWithReleasedField{
		Released: m.GetReleased(),
	}
	return msg
}
func (m *MessageWithReleasedField) FromBase(b *fixtures.MessageWithReleasedField) *MessageWithReleasedField {
	msg := &MessageWithReleasedField{
		Released: b.GetReleased(),
	}
	return msg
}
func (m *MessageWithReleasedOneofItem) ToBase() *fixtures.MessageWithReleasedOneofItem {
	msg := &fixtures.MessageWithReleasedOneofItem{}
	switch o := m.GetOneofWithItem().(type) {
	case *MessageWithReleasedOneofItem_OneofItemReleased:
		msg.OneofWithItem = o.ToBase()
	}
	return msg
}
func (m *MessageWithReleasedOneofItem) FromBase(b *fixtures.MessageWithReleasedOneofItem) *MessageWithReleasedOneofItem {
	msg := &MessageWithReleasedOneofItem{}
	switch o := b.GetOneofWithItem().(type) {
	case *fixtures.MessageWithReleasedOneofItem_OneofItemReleased:
		m.OneofWithItem = (*MessageWithReleasedOneofItem_OneofItemReleased)(nil).FromBase(o)
	}
	return msg
}
func (m *MessageWithReleasedOneofItem_OneofItemReleased) ToBase() *fixtures.MessageWithReleasedOneofItem_OneofItemReleased {
	return (*fixtures.MessageWithReleasedOneofItem_OneofItemReleased)(m)
}
func (m *MessageWithReleasedOneofItem_OneofItemReleased) FromBase(b *fixtures.MessageWithReleasedOneofItem_OneofItemReleased) *MessageWithReleasedOneofItem_OneofItemReleased {
	return (*MessageWithReleasedOneofItem_OneofItemReleased)(b)
}
func (m *MessageWithOneofWithMessages) ToBase() *fixtures.MessageWithOneofWithMessages {
	msg := &fixtures.MessageWithOneofWithMessages{}
	switch o := m.GetOneofWithMessage().(type) {
	case *MessageWithOneofWithMessages_MessageWithReleaseAnnotation:
		msg.OneofWithMessage = o.ToBase()
	}
	return msg
}
func (m *MessageWithOneofWithMessages) FromBase(b *fixtures.MessageWithOneofWithMessages) *MessageWithOneofWithMessages {
	msg := &MessageWithOneofWithMessages{}
	switch o := b.GetOneofWithMessage().(type) {
	case *fixtures.MessageWithOneofWithMessages_MessageWithReleaseAnnotation:
		m.OneofWithMessage = (*MessageWithOneofWithMessages_MessageWithReleaseAnnotation)(nil).FromBase(o)
	}
	return msg
}
func (m *MessageWithOneofWithMessages_MessageWithReleaseAnnotation) ToBase() *fixtures.MessageWithOneofWithMessages_MessageWithReleaseAnnotation {
	return &fixtures.MessageWithOneofWithMessages_MessageWithReleaseAnnotation{
		MessageWithReleaseAnnotation: m.MessageWithReleaseAnnotation.ToBase(),
	}
}
func (m *MessageWithOneofWithMessages_MessageWithReleaseAnnotation) FromBase(b *fixtures.MessageWithOneofWithMessages_MessageWithReleaseAnnotation) *MessageWithOneofWithMessages_MessageWithReleaseAnnotation {
	return &MessageWithOneofWithMessages_MessageWithReleaseAnnotation{
		MessageWithReleaseAnnotation: (*MessageWithReleasedField)(nil).FromBase(b.MessageWithReleaseAnnotation),
	}
}
func (m *MessageNotAnnotated) ToBase() *fixtures.MessageNotAnnotated {
	msg := &fixtures.MessageNotAnnotated{
		Released:              m.GetReleased(),
		PreviewedThenReleased: m.GetPreviewedThenReleased(),
	}
	switch o := m.GetNotAnnotatedOneof().(type) {
	case *MessageNotAnnotated_OneofItemNotAnnotated:
		msg.NotAnnotatedOneof = o.ToBase()
	}
	return msg
}
func (m *MessageNotAnnotated) FromBase(b *fixtures.MessageNotAnnotated) *MessageNotAnnotated {
	msg := &MessageNotAnnotated{
		Released:              b.GetReleased(),
		PreviewedThenReleased: b.GetPreviewedThenReleased(),
	}
	switch o := b.GetNotAnnotatedOneof().(type) {
	case *fixtures.MessageNotAnnotated_OneofItemNotAnnotated:
		m.NotAnnotatedOneof = (*MessageNotAnnotated_OneofItemNotAnnotated)(nil).FromBase(o)
	}
	return msg
}
func (m *MessageNotAnnotated_OneofItemNotAnnotated) ToBase() *fixtures.MessageNotAnnotated_OneofItemNotAnnotated {
	return (*fixtures.MessageNotAnnotated_OneofItemNotAnnotated)(m)
}
func (m *MessageNotAnnotated_OneofItemNotAnnotated) FromBase(b *fixtures.MessageNotAnnotated_OneofItemNotAnnotated) *MessageNotAnnotated_OneofItemNotAnnotated {
	return (*MessageNotAnnotated_OneofItemNotAnnotated)(b)
}
func (m *MessageWithEnumFields) ToBase() *fixtures.MessageWithEnumFields {
	msg := &fixtures.MessageWithEnumFields{
		Released: m.GetReleased().ToBase(),
	}
	return msg
}
func (m *MessageWithEnumFields) FromBase(b *fixtures.MessageWithEnumFields) *MessageWithEnumFields {
	msg := &MessageWithEnumFields{}
	msg.Released = msg.Released.FromBase(b.GetReleased())
	return msg
}
func (m *MessageWithOneofsWithEnumFields) ToBase() *fixtures.MessageWithOneofsWithEnumFields {
	msg := &fixtures.MessageWithOneofsWithEnumFields{}
	switch o := m.GetOneofWithEnumField().(type) {
	case *MessageWithOneofsWithEnumFields_Released:
		msg.OneofWithEnumField = o.ToBase()
	}
	return msg
}
func (m *MessageWithOneofsWithEnumFields) FromBase(b *fixtures.MessageWithOneofsWithEnumFields) *MessageWithOneofsWithEnumFields {
	msg := &MessageWithOneofsWithEnumFields{}
	switch o := b.GetOneofWithEnumField().(type) {
	case *fixtures.MessageWithOneofsWithEnumFields_Released:
		m.OneofWithEnumField = (*MessageWithOneofsWithEnumFields_Released)(nil).FromBase(o)
	}
	return msg
}
func (m *MessageWithOneofsWithEnumFields_Released) ToBase() *fixtures.MessageWithOneofsWithEnumFields_Released {
	return &fixtures.MessageWithOneofsWithEnumFields_Released{
		Released: m.Released.ToBase(),
	}
}
func (m *MessageWithOneofsWithEnumFields_Released) FromBase(b *fixtures.MessageWithOneofsWithEnumFields_Released) *MessageWithOneofsWithEnumFields_Released {
	return &MessageWithOneofsWithEnumFields_Released{
		Released: (EnumNotAnnotated)(0).FromBase(b.Released),
	}
}
func (e EnumNotAnnotated) FromBase(b fixtures.EnumNotAnnotated) EnumNotAnnotated {
	switch b {
	case fixtures.EnumNotAnnotated_zero_value:
		return EnumNotAnnotated_zero_value
	case fixtures.EnumNotAnnotated_released:
		return EnumNotAnnotated_released
	case fixtures.EnumNotAnnotated_previewed_then_released:
		return EnumNotAnnotated_previewed_then_released
	default:
		return EnumNotAnnotated(0)
	}
}
func (e EnumNotAnnotated) ToBase() fixtures.EnumNotAnnotated {
	return fixtures.EnumNotAnnotated(e)
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
