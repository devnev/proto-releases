// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

//go:generate sh -ec "[ ! -e releases ] || rm -r releases"
//go:generate sh -ec "[ ! -e server ] || rm -r server"
//go:generate mkdir -p releases/alpha releases/beta releases/stable
//go:generate go install github.com/devnev/proto-releases/protoc-gen-release
//go:generate -command gen-release protoc -I.. -I. example.proto
//go:generate gen-release --release_out=releases/alpha --release_opt=go_package=github.com/devnev/proto-releases/examples:releases/alpha --release_opt=package=example:alpha
//go:generate gen-release --release_out=releases/beta --release_opt=go_package=github.com/devnev/proto-releases/examples:releases/beta,release=3,preview=true --release_opt=package=example:beta
//go:generate gen-release --release_out=releases/stable --release_opt=go_package=github.com/devnev/proto-releases/examples:releases/stable,release=3 --release_opt=package=example:stable

//go:generate mkdir -p server/alpha server/beta server/stable
//go:generate go install github.com/devnev/proto-releases/protoc-gen-go-torelease
//go:generate protoc -I.. -I. --go_out=server --go_opt=paths=source_relative --go-grpc_out=server --go-grpc_opt=paths=source_relative --go-torelease_out=server --go-torelease_opt=paths=source_relative example.proto

//go:generate go install github.com/devnev/proto-releases/protoc-gen-go-baseconvert
//go:generate -command gen-server protoc -I.. -Ireleases --go_out=server --go_opt=paths=source_relative --go-grpc_out=server --go-grpc_opt=paths=source_relative --go-baseconvert_out=server --go-baseconvert_opt=paths=source_relative
//go:generate gen-server alpha/example.proto --go-baseconvert_opt=base_go_package=github.com/devnev/proto-releases/examples/releases/alpha:../../server
//go:generate gen-server beta/example.proto --go-baseconvert_opt=base_go_package=github.com/devnev/proto-releases/examples/releases/beta:../../server
//go:generate gen-server stable/example.proto --go-baseconvert_opt=base_go_package=github.com/devnev/proto-releases/examples/releases/stable:../../server
