// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package releases

import (
	"fmt"
	"strings"
	"unicode"
)

type PackageShorthand struct {
	Config    *Config_PackageMapping
	ConfigPtr **Config_PackageMapping
	setVal    string
}

type PackageShorthandError struct {
	shorthand    string
	invaildRune  rune
	invalidIndex int
}

func (err PackageShorthandError) Error() string {
	return fmt.Sprintf("invalid character %q at index %d of shorthand %q", err.invaildRune, err.invalidIndex, err.shorthand)
}

func (err PackageShorthandError) Unwrap() error {
	return ErrConfig
}

func (s *PackageShorthand) Set(v string) error {
	conf := s.Config
	if conf == nil && s.ConfigPtr != nil {
		conf = &Config_PackageMapping{}
	}
	if err := ParsePackageShorthand(v, s.Config); err != nil {
		return err
	}
	if s.ConfigPtr != nil {
		*s.ConfigPtr = conf
		s.Config = conf
	}
	s.setVal = v
	return nil
}

func (s *PackageShorthand) String() string {
	isSameConf := func(c1, c2 *Config_PackageMapping) bool {
		return c1.GetSourceRoot() == c2.GetSourceRoot() &&
			c1.GetReleaseRoot() == c2.GetReleaseRoot()
	}

	if s.setVal != "" {
		var tmpConf Config_PackageMapping
		if ParsePackageShorthand(s.setVal, &tmpConf) == nil && isSameConf(&tmpConf, s.Config) {
			return s.setVal
		}
	}

	val := formatShorthand(
		s.Config.GetSourceRoot(),
		s.Config.GetReleaseRoot(),
		s.Config.GetReleaseSuffix(),
		".",
	)
	if tmpConf := (&Config_PackageMapping{}); ParsePackageShorthand(val, tmpConf) != nil || !isSameConf(tmpConf, s.Config) {
		return "<invalid>" + val
	}
	return val
}

func ParsePackageShorthand(v string, into *Config_PackageMapping) error {
	into.Reset()

	parts, err := splitPackageShorthand(v, '.', isValidPackageRune)
	switch {
	case err != nil:
		return err
	case len(parts) == 1:
		into.ReleaseRoot = parts[0]
	case len(parts) == 3:
		into.ReleaseSuffix = parts[2]
		fallthrough
	case len(parts) == 2:
		into.SourceRoot = parts[0]
		into.ReleaseRoot = parts[1]
		if strings.HasPrefix(parts[1], ".") {
			into.ReleaseRoot = parts[0] + parts[1]
		}
	}

	return nil
}

func splitPackageShorthand(s string, pathSep rune, okRune func(rune) bool) ([]string, error) {
	var (
		parts        = make([]string, 0, 3)
		start        = 0
		afterPathSep = true
	)
	for i, c := range s {
		if c == ':' && len(parts) < 2 {
			parts = append(parts, s[start:i])
			start = i + 1
		} else if c == pathSep && afterPathSep {
			return nil, GoPackageShorthandError{s, c, i}
		} else if !okRune(c) {
			return nil, GoPackageShorthandError{s, c, i}
		}
		afterPathSep = c == pathSep
	}
	parts = append(parts, s[start:])
	return parts, nil
}

func formatShorthand(sourceRoot, releaseRoot, releaseSuffix, relativePrefix string) string {
	var val string
	if deprefixed := strings.TrimPrefix(releaseRoot, sourceRoot+"/"); len(deprefixed) < len(releaseRoot) {
		val = sourceRoot + ":" + relativePrefix + deprefixed
	} else if sourceRoot != "" {
		val = sourceRoot + ":" + releaseRoot
	} else {
		val = releaseRoot
	}
	if releaseSuffix != "" {
		val += ":" + releaseSuffix
	}
	return val
}

func isValidPackageRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r) || strings.ContainsRune("_.", r)
}
