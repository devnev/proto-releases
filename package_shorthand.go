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
	return ConfigError
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

	srcRoot, relRoot := s.Config.GetSourceRoot(), s.Config.GetReleaseRoot()
	var val string
	if deprefixed := strings.TrimPrefix(relRoot, srcRoot+"."); len(deprefixed) < len(relRoot) {
		val = srcRoot + ":" + deprefixed
	} else if srcRoot != "" {
		val = srcRoot + ";" + relRoot
	} else {
		val = relRoot
	}
	if tmpConf := (&Config_PackageMapping{}); ParsePackageShorthand(val, tmpConf) != nil || !isSameConf(tmpConf, s.Config) {
		return "<invalid>" + val
	}
	return val
}

func ParsePackageShorthand(v string, into *Config_PackageMapping) error {
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
		into.ReleaseRoot = into.SourceRoot + "." + v[idx+1:]
	}
	return nil
}

func validatePackageShorthand(s string) error {
	var sepCount int
	afterSep := true
	for i, c := range s {
		if c == ':' || c == ';' {
			sepCount++
			if sepCount > 1 {
				return PackageShorthandError{s, c, i}
			}
			afterSep = true
			continue
		}
		if afterSep && c == '.' {
			return PackageShorthandError{s, c, i}
		}
		if isInvalidGoPackageRune(c) {
			return PackageShorthandError{s, c, i}
		}
		afterSep = false
	}
	return nil
}

func isInvalidPackageRune(r rune) bool {
	return !(unicode.IsLetter(r) || unicode.IsNumber(r) || strings.ContainsRune("_.", r))
}
