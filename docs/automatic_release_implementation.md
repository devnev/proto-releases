# Automatic Release Implementation

The `protoc-gen-go-baseconvert` plugin can generate code to automatically
implement releases of GRPC services to use a single implementation of the
original proto definition.

This allows code like the following to serve both the base service and the
subset of the service available in a particular release without duplicate
implementations.

```go
func main() {
    srv := grpc.NewServer()
    var apiSrv server.ApiServer = &myServerImpl{}
    server.RegisterApiServer(srv, apiSrv)
    beta17.RegisterApiBaseServer(srv, apiSrv)
}
```

To generate the release implementations, run something like the following:

```sh
# Install tools. for installing protoc, see https://grpc.io/docs/protoc-installation/.
go install \
    github.com/devnev/proto-releases/protoc-gen-release \
    github.com/devnev/proto-releases/protoc-gen-go-baseconvert \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    google.golang.org/protobuf/cmd/protoc-gen-go

# Generate the preview release 17 ("beta17") of my_api.proto.
# Due to the go_package option, release proto's go_package will start with
# github.com/me/releases/beta17, followed by the original go_package but with
# the prefix github.com/me stripped.
protoc my_api.proto \
    --release_out=releases/beta17 \
    --release_opt=release=17,preview=true,go_package=github.com/me:releases/beta17

# Generate internal protos and server stubs for my_api.proto.
# In this example, stubs are being generated in a separate "server" directory.
protoc my_api.proto \
    --go_out=server --go_opt=paths=source_relative \
    --go-grpc_out=server --go-grpc_opt=paths=source_relative

# Generate automatic implementation of the beta17 release.
protoc my_api.proto \
    -Ireleases/beta17 \
    --go_out=server/beta17 --go_opt=paths=source_relative \
    --go-grpc_out=server/beta17 --go-grpc_opt=paths=source_relative \
    --go-baseconvert_out=server/beta17 \
    --go-baseconvert_opt=base_go_package=github.com/me:server\
    --go-baseconvert_opt=paths=source_relative
```
