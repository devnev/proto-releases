// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package transform

import (
	"strconv"
	"testing"

	releases "github.com/devnev/proto-releases"
)

func TestGoPackage(t *testing.T) {
	for testIdx, tc := range []struct {
		p string
		i string
		o string
	}{
		{
			"",
			"a/b",
			"a/b",
		},
		{
			"a:c",
			"a/b",
			"c/b",
		},
		{
			"a:./c",
			"a/b",
			"a/c/b",
		},
		{
			"a:../c",
			"a/b",
			"c/b",
		},
		{
			"::c",
			"a/b",
			"a/b/c",
		},
	} {
		t.Run(strconv.Itoa(testIdx), func(t *testing.T) {
			t.Logf("%#v", tc)
			var c releases.Config_GoPackageMapping
			if err := releases.ParseGoPackageShorthand(tc.p, &c); err != nil {
				t.Errorf("failed to parse %q: %v", tc.p, err)
			}
			t.Log(c.String())
			o := GoPackage(tc.i, &c)
			if tc.o != o {
				t.Errorf("expected %q to transform %q to %q, got %q", tc.p, tc.i, tc.o, o)
			}
		})
	}
}

func TestPackage(t *testing.T) {
	for testIdx, tc := range []struct {
		p string
		i string
		o string
	}{
		{
			"",
			"a.b",
			"a.b",
		},
		{
			"a:c",
			"a.b",
			"c.b",
		},
		{
			"a:.c",
			"a.b",
			"a.c.b",
		},
		{
			"::c",
			"a.b",
			"a.b.c",
		},
	} {
		t.Run(strconv.Itoa(testIdx), func(t *testing.T) {
			t.Logf("%#v", tc)
			var c releases.Config_PackageMapping
			if err := releases.ParsePackageShorthand(tc.p, &c); err != nil {
				t.Errorf("failed to parse %q: %v", tc.p, err)
			}
			t.Log(c.String())
			o := Package(tc.i, &c)
			if tc.o != o {
				t.Errorf("expected %q to transform %q to %q, got %q", tc.p, tc.i, tc.o, o)
			}
		})
	}
}
