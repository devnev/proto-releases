package releases

import (
	"fmt"
	"path"
	"strings"
	"unicode"
)

type GoPackageFlagValue struct {
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
	return fmt.Sprintf("invalid character %q at byte %d of shorthand %q", err.invaildRune, err.invalidIndex, err.shorthand)
}

func (err GoPackageShorthandError) Unwrap() error {
	return ConfigError
}

func (s *GoPackageFlagValue) Set(v string) error {
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

func (s *GoPackageFlagValue) String() string {
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
		val = srcRoot + ":./" + deprefixed
	} else if srcRoot != "" {
		val = srcRoot + ":" + relRoot
	} else {
		val = relRoot
	}
	if s.Config.GetReleaseSuffix() != "" {
		val += ":" + s.Config.GetReleaseSuffix()
	}
	if tmpConf := (&Config_GoPackageMapping{}); ParseGoPackageShorthand(val, tmpConf) != nil || !isSameConf(tmpConf, s.Config) {
		return "<invalid>" + val
	}
	return val
}

func ParseGoPackageShorthand(v string, into *Config_GoPackageMapping) error {
	into.Reset()

	parts, err := splitPackageShorthand(v, '/', isValidGoPackageRune)
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
		if strings.HasPrefix(parts[1], "./") || strings.HasPrefix(parts[1], "../") {
			into.ReleaseRoot = path.Join(parts[0], parts[1])
		}
	}

	return nil
}

func isValidGoPackageRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r) || strings.ContainsRune("-_./", r)
}
