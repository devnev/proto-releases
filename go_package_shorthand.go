package releases

import (
	"fmt"
	"path"
	"strings"
	"unicode"
)

type GoPackageShorthand struct {
	Config    *Config_GoPackageMapping
	ConfigPtr **Config_GoPackageMapping
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
		conf = &Config_GoPackageMapping{}
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
	isSameConf := func(c1, c2 *Config_GoPackageMapping) bool {
		return c1.GetSourceRoot() == c2.GetSourceRoot() &&
			c1.GetReleaseRoot() == c2.GetReleaseRoot()
	}

	if s.setVal != "" {
		var tmpConf Config_GoPackageMapping
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
	if tmpConf := (&Config_GoPackageMapping{}); ParseGoPackageShorthand(val, tmpConf) != nil || !isSameConf(tmpConf, s.Config) {
		return "<invalid>" + val
	}
	return val
}

func ParseGoPackageShorthand(v string, into *Config_GoPackageMapping) error {
	if err := validateGoPackageShorthand(v); err != nil {
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

func validateGoPackageShorthand(s string) error {
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
		if isInvalidGoPackageRune(c) {
			return GoPackageShorthandError{s, c, i}
		}
		afterSep = false
	}
	return nil
}

func isInvalidGoPackageRune(r rune) bool {
	return !(unicode.IsLetter(r) || unicode.IsNumber(r) || strings.ContainsRune("/-_.", r))
}
