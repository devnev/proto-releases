// Code generated by protoc-gen-go-baseconvert. DO NOT EDIT.

package fixtures

import (
	fixtures "github.com/devnev/proto-releases/fixtures"
)

func (m *ImportedNotAnnotatedWithReleasedField) ToBase() *fixtures.ImportedNotAnnotatedWithReleasedField {
	msg := &fixtures.ImportedNotAnnotatedWithReleasedField{
		Released: m.GetReleased(),
	}
	return msg
}
func (m *ImportedNotAnnotatedWithReleasedField) FromBase(b *fixtures.ImportedNotAnnotatedWithReleasedField) *ImportedNotAnnotatedWithReleasedField {
	msg := &ImportedNotAnnotatedWithReleasedField{
		Released: b.GetReleased(),
	}
	return msg
}
