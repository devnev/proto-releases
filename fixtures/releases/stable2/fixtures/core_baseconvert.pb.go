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
	if m != nil {
		m.Reset()
	} else {
		m = new(EmptyMessageReleased)
	}
	return m
}
func (m *EmptyMessageReleasedThenRemoved) ToBase() *fixtures.EmptyMessageReleasedThenRemoved {
	msg := &fixtures.EmptyMessageReleasedThenRemoved{}
	return msg
}
func (m *EmptyMessageReleasedThenRemoved) FromBase(b *fixtures.EmptyMessageReleasedThenRemoved) *EmptyMessageReleasedThenRemoved {
	if m != nil {
		m.Reset()
	} else {
		m = new(EmptyMessageReleasedThenRemoved)
	}
	return m
}
func (m *MessageWithReleasedField) ToBase() *fixtures.MessageWithReleasedField {
	msg := &fixtures.MessageWithReleasedField{
		Released: m.GetReleased(),
	}
	return msg
}
func (m *MessageWithReleasedField) FromBase(b *fixtures.MessageWithReleasedField) *MessageWithReleasedField {
	if m != nil {
		m.Reset()
	} else {
		m = new(MessageWithReleasedField)
	}
	m.Released = b.GetReleased()
	return m
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
	if m != nil {
		m.Reset()
	} else {
		m = new(MessageWithReleasedOneofItem)
	}
	switch o := b.GetOneofWithItem().(type) {
	case *fixtures.MessageWithReleasedOneofItem_OneofItemReleased:
		m.OneofWithItem = (*MessageWithReleasedOneofItem_OneofItemReleased)(nil).FromBase(o)
	}
	return m
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
	if m != nil {
		m.Reset()
	} else {
		m = new(MessageWithOneofWithMessages)
	}
	switch o := b.GetOneofWithMessage().(type) {
	case *fixtures.MessageWithOneofWithMessages_MessageWithReleaseAnnotation:
		m.OneofWithMessage = (*MessageWithOneofWithMessages_MessageWithReleaseAnnotation)(nil).FromBase(o)
	}
	return m
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
func (m *MessageWithImportedFields) ToBase() *fixtures.MessageWithImportedFields {
	msg := &fixtures.MessageWithImportedFields{
		EmptyReleased: m.GetEmptyReleased().ToBase(),
	}
	return msg
}
func (m *MessageWithImportedFields) FromBase(b *fixtures.MessageWithImportedFields) *MessageWithImportedFields {
	if m != nil {
		m.Reset()
	} else {
		m = new(MessageWithImportedFields)
	}
	m.EmptyReleased = m.EmptyReleased.FromBase(b.GetEmptyReleased())
	return m
}
func (m *MessageNotAnnotated) ToBase() *fixtures.MessageNotAnnotated {
	msg := &fixtures.MessageNotAnnotated{
		Released:            m.GetReleased(),
		ReleasedThenRemoved: m.GetReleasedThenRemoved(),
	}
	switch o := m.GetNotAnnotatedOneof().(type) {
	case *MessageNotAnnotated_OneofItemNotAnnotated:
		msg.NotAnnotatedOneof = o.ToBase()
	}
	return msg
}
func (m *MessageNotAnnotated) FromBase(b *fixtures.MessageNotAnnotated) *MessageNotAnnotated {
	if m != nil {
		m.Reset()
	} else {
		m = new(MessageNotAnnotated)
	}
	m.Released = b.GetReleased()
	m.ReleasedThenRemoved = b.GetReleasedThenRemoved()
	switch o := b.GetNotAnnotatedOneof().(type) {
	case *fixtures.MessageNotAnnotated_OneofItemNotAnnotated:
		m.NotAnnotatedOneof = (*MessageNotAnnotated_OneofItemNotAnnotated)(nil).FromBase(o)
	}
	return m
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
	if m != nil {
		m.Reset()
	} else {
		m = new(MessageWithEnumFields)
	}
	m.Released = m.Released.FromBase(b.GetReleased())
	return m
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
	if m != nil {
		m.Reset()
	} else {
		m = new(MessageWithOneofsWithEnumFields)
	}
	switch o := b.GetOneofWithEnumField().(type) {
	case *fixtures.MessageWithOneofsWithEnumFields_Released:
		m.OneofWithEnumField = (*MessageWithOneofsWithEnumFields_Released)(nil).FromBase(o)
	}
	return m
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
	case fixtures.EnumNotAnnotated_released_then_removed:
		return EnumNotAnnotated_released_then_removed
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
