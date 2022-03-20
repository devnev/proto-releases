package releases

import (
	"errors"
	"fmt"
	"path"
	"strings"
	"unicode"
)

var ConfigError = errors.New("release config error")

type GoPackageShorthand struct {
	Config    *Config_GoPackage
	ConfigPtr **Config_GoPackage
	setVal    string
}

type GoPackageShorthandError struct {
	shorthand    string
	invaildRune  rune
	invalidIndex int
}

func (err GoPackageShorthandError) Error() string {
	return fmt.Sprintf("invalid character %q at index %d of shorthand %q", err.invaildRune, err.invalidIndex, err.shorthand)
}

func (err GoPackageShorthandError) Unwrap() error {
	return ConfigError
}

func (s *GoPackageShorthand) Set(v string) error {
	conf := s.Config
	if conf == nil && s.ConfigPtr != nil {
		conf = &Config_GoPackage{}
	}
	if err := ParseGoPackageShorthand(v, s.Config); err != nil {
		return err
	}
	if s.ConfigPtr != nil {
		*s.ConfigPtr = conf
		s.Config = conf
	}
	s.setVal = v
	return nil
}

func (s *GoPackageShorthand) String() string {
	isSameConf := func(c1, c2 *Config_GoPackage) bool {
		return c1.GetSourceRoot() == c2.GetSourceRoot() &&
			c1.GetReleaseRoot() == c2.GetReleaseRoot()
	}

	if s.setVal != "" {
		var tmpConf Config_GoPackage
		if ParseGoPackageShorthand(s.setVal, &tmpConf) == nil && isSameConf(&tmpConf, s.Config) {
			return s.setVal
		}
	}

	srcRoot, relRoot := s.Config.GetSourceRoot(), s.Config.GetReleaseRoot()
	var val string
	if deprefixed := strings.TrimPrefix(relRoot, srcRoot+"/"); len(deprefixed) < len(relRoot) {
		val = srcRoot + ":" + deprefixed
	} else if srcRoot != "" {
		val = srcRoot + ";" + relRoot
	} else {
		val = relRoot
	}
	if tmpConf := (&Config_GoPackage{}); ParseGoPackageShorthand(val, tmpConf) != nil || !isSameConf(tmpConf, s.Config) {
		return "<invalid>" + val
	}
	return val
}

func ParseGoPackageShorthand(v string, into *Config_GoPackage) error {
	if err := validateShorthand(v); err != nil {
		return err
	}
	if idx := strings.Index(v, ";"); idx >= 0 {
		into.SourceRoot = v[:idx]
		into.ReleaseRoot = v[idx+1:]
	} else if idx := strings.Index(v, ":"); idx < 0 {
		into.SourceRoot = ""
		into.ReleaseRoot = v
	} else {
		into.SourceRoot = v[:idx]
		into.ReleaseRoot = path.Join(into.SourceRoot, v[idx+1:])
	}
	return nil
}

func validateShorthand(s string) error {
	var sepCount int
	afterSep := true
	for i, c := range s {
		if c == ':' || c == ';' {
			sepCount++
			if sepCount > 1 {
				return GoPackageShorthandError{s, c, i}
			}
			afterSep = true
			continue
		}
		if afterSep && c == '/' {
			return GoPackageShorthandError{s, c, i}
		}
		if isInvalidPackageRune(c) {
			return GoPackageShorthandError{s, c, i}
		}
		afterSep = false
	}
	return nil
}

func isInvalidPackageRune(r rune) bool {
	return !(unicode.IsLetter(r) || unicode.IsNumber(r) || strings.ContainsRune("/-_.", r))
}
