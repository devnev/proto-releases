// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package releases

//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative releases.proto config.proto
//go:generate mkdir -p testdata/golden/alpha testdata/golden/beta1 testdata/golden/beta2 testdata/golden/beta3 testdata/golden/stable1 testdata/golden/stable2 testdata/golden/stable3
//go:generate go install github.com/devnev/proto-releases/protoc-gen-release
//go:generate -command genrel protoc -I. -Itestdata release-option-combinations.proto
//go:generate genrel --release_out=testdata/golden/alpha --release_opt=go_package=github.com/devnev/proto-releases/alpha
//go:generate genrel --release_out=testdata/golden/beta1 --release_opt=go_package=github.com/devnev/proto-releases/beta1,release=1,preview=true
//go:generate genrel --release_out=testdata/golden/beta2 --release_opt=go_package=github.com/devnev/proto-releases/beta2,release=2,preview=true
//go:generate genrel --release_out=testdata/golden/beta3 --release_opt=go_package=github.com/devnev/proto-releases/beta3,release=3,preview=true
//go:generate genrel --release_out=testdata/golden/stable1 --release_opt=go_package=github.com/devnev/proto-releases/stable1,release=1
//go:generate genrel --release_out=testdata/golden/stable2 --release_opt=go_package=github.com/devnev/proto-releases/stable2,release=2
//go:generate genrel --release_out=testdata/golden/stable3 --release_opt=go_package=github.com/devnev/proto-releases/stable3,release=3
//go:generate protoc -I. -Itestdata --go_out=testdata/golden --go_opt=paths=source_relative release-option-combinations.proto
//go:generate go install github.com/devnev/proto-releases/protoc-gen-go-baseconvert
//go:generate -command genrelgo protoc -I. -Itestdata/golden --go_out=testdata/golden --go_opt=paths=source_relative --go-grpc_out=testdata/golden --go-grpc_opt=paths=source_relative --go-baseconvert_out=testdata/golden --go-baseconvert_opt=base=github.com/devnev/proto-releases/testdata/golden,paths=source_relative
//go:generate genrelgo alpha/release-option-combinations.proto
//go:generate genrelgo beta1/release-option-combinations.proto
//go:generate genrelgo beta2/release-option-combinations.proto
//go:generate genrelgo beta3/release-option-combinations.proto
//go:generate genrelgo stable1/release-option-combinations.proto
//go:generate genrelgo stable2/release-option-combinations.proto
//go:generate genrelgo stable3/release-option-combinations.proto
