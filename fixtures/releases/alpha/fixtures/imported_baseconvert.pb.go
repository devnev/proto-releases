// Code generated by protoc-gen-go-baseconvert. DO NOT EDIT.

package fixtures

import (
	fixtures "github.com/devnev/proto-releases/fixtures"
)

func (m *ImportedNotAnnotatedAndEmpty) ToBase() *fixtures.ImportedNotAnnotatedAndEmpty {
	msg := &fixtures.ImportedNotAnnotatedAndEmpty{}
	return msg
}
func (m *ImportedNotAnnotatedAndEmpty) FromBase(b *fixtures.ImportedNotAnnotatedAndEmpty) *ImportedNotAnnotatedAndEmpty {
	msg := &ImportedNotAnnotatedAndEmpty{}
	return msg
}
func (m *ImportedNotAnnotatedWithUnreleasedField) ToBase() *fixtures.ImportedNotAnnotatedWithUnreleasedField {
	msg := &fixtures.ImportedNotAnnotatedWithUnreleasedField{
		Released: m.GetReleased(),
	}
	return msg
}
func (m *ImportedNotAnnotatedWithUnreleasedField) FromBase(b *fixtures.ImportedNotAnnotatedWithUnreleasedField) *ImportedNotAnnotatedWithUnreleasedField {
	msg := &ImportedNotAnnotatedWithUnreleasedField{
		Released: b.GetReleased(),
	}
	return msg
}
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