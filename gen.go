// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package releases

//go:generate rm -r fixtures/releases
//go:generate find fixtures -name *.go -delete
//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative releases.proto config.proto
//go:generate mkdir -p fixtures/releases/alpha fixtures/releases/beta1 fixtures/releases/beta2 fixtures/releases/beta3 fixtures/releases/stable1 fixtures/releases/stable2 fixtures/releases/stable3
//go:generate go install github.com/devnev/proto-releases/protoc-gen-release
//go:generate -command genrel protoc -I. fixtures/core.proto fixtures/imported.proto fixtures/subpackage/subimport.proto releases.proto
//go:generate genrel --release_out=fixtures/releases/alpha --release_opt=go_package=github.com/devnev/proto-releases:fixtures/releases/alpha
//go:generate genrel --release_out=fixtures/releases/beta1 --release_opt=go_package=github.com/devnev/proto-releases:fixtures/releases/beta1,release=1,preview=true
//go:generate genrel --release_out=fixtures/releases/beta2 --release_opt=go_package=github.com/devnev/proto-releases:fixtures/releases/beta2,release=2,preview=true
//go:generate genrel --release_out=fixtures/releases/beta3 --release_opt=go_package=github.com/devnev/proto-releases:fixtures/releases/beta3,release=3,preview=true
//go:generate genrel --release_out=fixtures/releases/stable1 --release_opt=go_package=github.com/devnev/proto-releases:fixtures/releases/stable1,release=1
//go:generate genrel --release_out=fixtures/releases/stable2 --release_opt=go_package=github.com/devnev/proto-releases:fixtures/releases/stable2,release=2
//go:generate genrel --release_out=fixtures/releases/stable3 --release_opt=go_package=github.com/devnev/proto-releases:fixtures/releases/stable3,release=3
//go:generate protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative fixtures/core.proto fixtures/imported.proto fixtures/subpackage/subimport.proto

// TODO: stop including release.proto in releases once it is getting properly pruned from imports.
//go:generate -command genrelprtgo protoc releases.proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
//go:generate genrelprtgo -Ifixtures/releases/alpha --go_out=fixtures/releases/alpha --go-grpc_out=fixtures/releases/alpha
//go:generate genrelprtgo -Ifixtures/releases/beta1 --go_out=fixtures/releases/beta1 --go-grpc_out=fixtures/releases/beta1
//go:generate genrelprtgo -Ifixtures/releases/beta2 --go_out=fixtures/releases/beta2 --go-grpc_out=fixtures/releases/beta2
//go:generate genrelprtgo -Ifixtures/releases/beta3 --go_out=fixtures/releases/beta3 --go-grpc_out=fixtures/releases/beta3
//go:generate genrelprtgo -Ifixtures/releases/stable1 --go_out=fixtures/releases/stable1 --go-grpc_out=fixtures/releases/stable1
//go:generate genrelprtgo -Ifixtures/releases/stable2 --go_out=fixtures/releases/stable2 --go-grpc_out=fixtures/releases/stable2
//go:generate genrelprtgo -Ifixtures/releases/stable3 --go_out=fixtures/releases/stable3 --go-grpc_out=fixtures/releases/stable3

//go:generate go install github.com/devnev/proto-releases/protoc-gen-go-baseconvert
//go:generate -command genrelgo protoc fixtures/core.proto fixtures/imported.proto fixtures/subpackage/subimport.proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-baseconvert_opt=base_path=.,paths=source_relative
//go:generate genrelgo -Ifixtures/releases/alpha --go_out=fixtures/releases/alpha --go-grpc_out=fixtures/releases/alpha --go-baseconvert_out=fixtures/releases/alpha
//go:generate genrelgo -Ifixtures/releases/beta1 --go_out=fixtures/releases/beta1 --go-grpc_out=fixtures/releases/beta1 --go-baseconvert_out=fixtures/releases/beta1
//go:generate genrelgo -Ifixtures/releases/beta2 --go_out=fixtures/releases/beta2 --go-grpc_out=fixtures/releases/beta2 --go-baseconvert_out=fixtures/releases/beta2
//go:generate genrelgo -Ifixtures/releases/beta3 --go_out=fixtures/releases/beta3 --go-grpc_out=fixtures/releases/beta3 --go-baseconvert_out=fixtures/releases/beta3
//go:generate genrelgo -Ifixtures/releases/stable1 --go_out=fixtures/releases/stable1 --go-grpc_out=fixtures/releases/stable1 --go-baseconvert_out=fixtures/releases/stable1
//go:generate genrelgo -Ifixtures/releases/stable2 --go_out=fixtures/releases/stable2 --go-grpc_out=fixtures/releases/stable2 --go-baseconvert_out=fixtures/releases/stable2
//go:generate genrelgo -Ifixtures/releases/stable3 --go_out=fixtures/releases/stable3 --go-grpc_out=fixtures/releases/stable3 --go-baseconvert_out=fixtures/releases/stable3
