// Code generated by protoc-gen-go-baseconvert. DO NOT EDIT.

package subpackage

import (
	subpackage "github.com/devnev/proto-releases/fixtures/subpackage"
)

func (m *ImportedNotAnnotatedWithReleasedField) ToBase() *subpackage.ImportedNotAnnotatedWithReleasedField {
	msg := &subpackage.ImportedNotAnnotatedWithReleasedField{
		Released: m.GetReleased(),
	}
	return msg
}
func (m *ImportedNotAnnotatedWithReleasedField) FromBase(b *subpackage.ImportedNotAnnotatedWithReleasedField) *ImportedNotAnnotatedWithReleasedField {
	if m != nil {
		m.Reset()
	} else {
		m = new(ImportedNotAnnotatedWithReleasedField)
	}
	m.Released = b.GetReleased()
	return m
}
