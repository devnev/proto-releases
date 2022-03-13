// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package releases

//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative releases.proto config.proto
//go:generate mkdir -p fixtures/releases/alpha fixtures/releases/beta1 fixtures/releases/beta2 fixtures/releases/beta3 fixtures/releases/stable1 fixtures/releases/stable2 fixtures/releases/stable3
//go:generate go install github.com/devnev/proto-releases/protoc-gen-release
//go:generate -command genrel protoc -I. -Ifixtures release-option-combinations.proto
//go:generate genrel --release_out=fixtures/releases/alpha --release_opt=go_package=github.com/devnev/proto-releases/alpha
//go:generate genrel --release_out=fixtures/releases/beta1 --release_opt=go_package=github.com/devnev/proto-releases/beta1,release=1,preview=true
//go:generate genrel --release_out=fixtures/releases/beta2 --release_opt=go_package=github.com/devnev/proto-releases/beta2,release=2,preview=true
//go:generate genrel --release_out=fixtures/releases/beta3 --release_opt=go_package=github.com/devnev/proto-releases/beta3,release=3,preview=true
//go:generate genrel --release_out=fixtures/releases/stable1 --release_opt=go_package=github.com/devnev/proto-releases/stable1,release=1
//go:generate genrel --release_out=fixtures/releases/stable2 --release_opt=go_package=github.com/devnev/proto-releases/stable2,release=2
//go:generate genrel --release_out=fixtures/releases/stable3 --release_opt=go_package=github.com/devnev/proto-releases/stable3,release=3
//go:generate protoc -I. -Ifixtures --go_out=fixtures --go_opt=paths=source_relative --go-grpc_out=fixtures --go-grpc_opt=paths=source_relative release-option-combinations.proto
//go:generate go install github.com/devnev/proto-releases/protoc-gen-go-baseconvert
//go:generate -command genrelgo protoc -I. -Ifixtures/releases --go_out=fixtures/releases --go_opt=paths=source_relative --go-grpc_out=fixtures/releases --go-grpc_opt=paths=source_relative --go-baseconvert_out=fixtures/releases --go-baseconvert_opt=base=github.com/devnev/proto-releases/fixtures,paths=source_relative
//go:generate genrelgo alpha/release-option-combinations.proto
//go:generate genrelgo beta1/release-option-combinations.proto
//go:generate genrelgo beta2/release-option-combinations.proto
//go:generate genrelgo beta3/release-option-combinations.proto
//go:generate genrelgo stable1/release-option-combinations.proto
//go:generate genrelgo stable2/release-option-combinations.proto
//go:generate genrelgo stable3/release-option-combinations.proto
