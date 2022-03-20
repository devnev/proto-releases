// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package example

//go:generate sh -ec "[ ! -e releases ] || rm -r releases"
//go:generate sh -ec "[ ! -e server ] || rm -r server"
//go:generate mkdir -p releases/alpha releases/beta1 releases/beta2 releases/beta3 releases/stable1 releases/stable2 releases/stable3
//go:generate go install github.com/devnev/proto-releases/protoc-gen-release
//go:generate -command gen-release protoc -I.. -I. example.proto
//go:generate gen-release --release_out=releases/alpha --release_opt=go_package=github.com/devnev/proto-releases/examples/releases/alpha
//go:generate gen-release --release_out=releases/beta1 --release_opt=go_package=github.com/devnev/proto-releases/examples/releases/beta1,release=1,preview=true
//go:generate gen-release --release_out=releases/beta2 --release_opt=go_package=github.com/devnev/proto-releases/examples/releases/beta2,release=2,preview=true
//go:generate gen-release --release_out=releases/beta3 --release_opt=go_package=github.com/devnev/proto-releases/examples/releases/beta3,release=3,preview=true
//go:generate gen-release --release_out=releases/stable1 --release_opt=go_package=github.com/devnev/proto-releases/examples/releases/stable1,release=1
//go:generate gen-release --release_out=releases/stable2 --release_opt=go_package=github.com/devnev/proto-releases/examples/releases/stable2,release=2
//go:generate gen-release --release_out=releases/stable3 --release_opt=go_package=github.com/devnev/proto-releases/examples/releases/stable3,release=3

//go:generate mkdir -p server/alpha server/beta1 server/beta2 server/beta3 server/stable1 server/stable2 server/stable3
//go:generate protoc -I.. -I. --go_out=server --go_opt=paths=source_relative --go-grpc_out=server --go-grpc_opt=paths=source_relative example.proto
//go:generate go install github.com/devnev/proto-releases/protoc-gen-go-baseconvert
//go:generate -command gen-server protoc -I.. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-baseconvert_opt=base_path=.,base_go_package=github.com/devnev/proto-releases/examples:server,paths=source_relative example.proto
//go:generate gen-server -Ireleases/alpha --go_out=server/alpha --go-grpc_out=server/alpha --go-baseconvert_out=server/alpha
//go:generate gen-server -Ireleases/beta1 --go_out=server/beta1 --go-grpc_out=server/beta1 --go-baseconvert_out=server/beta1
//go:generate gen-server -Ireleases/beta2 --go_out=server/beta2 --go-grpc_out=server/beta2 --go-baseconvert_out=server/beta2
//go:generate gen-server -Ireleases/beta3 --go_out=server/beta3 --go-grpc_out=server/beta3 --go-baseconvert_out=server/beta3
//go:generate gen-server -Ireleases/stable1 --go_out=server/stable1 --go-grpc_out=server/stable1 --go-baseconvert_out=server/stable1
//go:generate gen-server -Ireleases/stable2 --go_out=server/stable2 --go-grpc_out=server/stable2 --go-baseconvert_out=server/stable2
//go:generate gen-server -Ireleases/stable3 --go_out=server/stable3 --go-grpc_out=server/stable3 --go-baseconvert_out=server/stable3
