// Code generated by protoc-gen-go-torelease. DO NOT EDIT.

package fixtures

import (
	proto_releases "github.com/devnev/proto-releases"
)

func (m *EmptyMessageNotAnnotated) ToRelease(c *proto_releases.Config) {
}
func (m *EmptyMessageReleased) ToRelease(c *proto_releases.Config) {
}
func (m *EmptyMessagePreviewed) ToRelease(c *proto_releases.Config) {
}
func (m *EmptyMessagePreviewedThenReleased) ToRelease(c *proto_releases.Config) {
}
func (m *EmptyMessagePreviewedThenRemoved) ToRelease(c *proto_releases.Config) {
}
func (m *EmptyMessageReleasedThenRemoved) ToRelease(c *proto_releases.Config) {
}
func (m *MessageWithNoAnnotations) ToRelease(c *proto_releases.Config) {
	if m == nil || c.GetRelease() == 0 {
		return
	}
	m.NotAnnotated = 0
}
func (m *MessageWithReleasedField) ToRelease(c *proto_releases.Config) {
	if m == nil || c.GetRelease() == 0 {
		return
	}
	r, p := c.GetRelease(), c.GetPreview()
	_, _ = r, p // prevent unused variable
	if !(r >= 2) {
		m.Released = 0
	}
}
func (m *MessageWithPreviewField) ToRelease(c *proto_releases.Config) {
	if m == nil || c.GetRelease() == 0 {
		return
	}
	r, p := c.GetRelease(), c.GetPreview()
	_, _ = r, p // prevent unused variable
	if !(p && r >= 2) {
		m.Released = 0
	}
}
func (m *MessageWithUnannotatedOneof) ToRelease(c *proto_releases.Config) {
	if m == nil || c.GetRelease() == 0 {
		return
	}
	switch o := m.NotAnnotatedOneof.(type) {
	default:
		_ = o // prevent unused variable
		m.NotAnnotatedOneof = nil
	}
}
func (m *MessageWithReleasedOneofItem) ToRelease(c *proto_releases.Config) {
	if m == nil || c.GetRelease() == 0 {
		return
	}
	r, p := c.GetRelease(), c.GetPreview()
	_, _ = r, p // prevent unused variable
	switch o := m.OneofWithItem.(type) {
	case *MessageWithReleasedOneofItem_ReleasedOneofItem:
		if !(r >= 2) {
			m.OneofWithItem = nil
		}
	default:
		_ = o // prevent unused variable
		m.OneofWithItem = nil
	}
}
func (m *MessageWithOneofWithMessages) ToRelease(c *proto_releases.Config) {
	if m == nil || c.GetRelease() == 0 {
		return
	}
	r, p := c.GetRelease(), c.GetPreview()
	_, _ = r, p // prevent unused variable
	switch o := m.OneofWithMessage.(type) {
	case *MessageWithOneofWithMessages_MessageWithNoAnnotations:
		o.MessageWithNoAnnotations.ToRelease(c)
	case *MessageWithOneofWithMessages_MessageWithReleasedField:
		o.MessageWithReleasedField.ToRelease(c)
	case *MessageWithOneofWithMessages_MessageWithReleaseAnnotation:
		if !(r >= 2) {
			m.OneofWithMessage = nil
		} else {
			o.MessageWithReleaseAnnotation.ToRelease(c)
		}
	default:
		_ = o // prevent unused variable
		m.OneofWithMessage = nil
	}
}
func (m *MessageWithImportedFields) ToRelease(c *proto_releases.Config) {
	if m == nil || c.GetRelease() == 0 {
		return
	}
	m.Empty.ToRelease(c)
	m.WithUnreleased.ToRelease(c)
	m.WithReleased.ToRelease(c)
}
func (m *MessageNotAnnotated) ToRelease(c *proto_releases.Config) {
	if m == nil || c.GetRelease() == 0 {
		return
	}
	r, p := c.GetRelease(), c.GetPreview()
	_, _ = r, p // prevent unused variable
	m.NotAnnotated = 0
	if !(r >= 2) {
		m.Released = 0
	}
	if !(p && r >= 2) {
		m.Previewed = 0
	}
	if !(r >= 3 || (p && r >= 2)) {
		m.PreviewedThenReleased = 0
	}
	if !(r < 3 && (p && r >= 2)) {
		m.PreviewedThenRemoved = 0
	}
	if !(r < 3 && (r >= 2)) {
		m.ReleasedThenRemoved = 0
	}
	switch o := m.NotAnnotatedOneof.(type) {
	case *MessageNotAnnotated_OneofItemNotAnnotated:
		if !(r >= 2) {
			m.NotAnnotatedOneof = nil
		}
	default:
		_ = o // prevent unused variable
		m.NotAnnotatedOneof = nil
	}
}
