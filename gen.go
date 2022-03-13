// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package releases

//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative releases.proto config.proto
//go:generate -command gen go run github.com/devnev/proto-releases/gen-release -inc .:testdata
//go:generate gen -out testdata/golden/alpha release-option-combinations.proto
//go:generate gen -out testdata/golden/beta1 -rel 1 -pre release-option-combinations.proto
//go:generate gen -out testdata/golden/beta2 -rel 2 -pre release-option-combinations.proto
//go:generate gen -out testdata/golden/beta3 -rel 3 -pre release-option-combinations.proto
//go:generate gen -out testdata/golden/stable1 -rel 1 release-option-combinations.proto
//go:generate gen -out testdata/golden/stable2 -rel 2 release-option-combinations.proto
//go:generate gen -out testdata/golden/stable3 -rel 3 release-option-combinations.proto
