# Proto-Releases

Proto-releases aims to provide protobuf extensions and tools to manage releases
of proto definitions. It currently allows creating unstable/alpha, preview/beta
and stable release channels of the same proto definitions.

This is experimental, work in progress, and by no means a complete set of tools
for the above purpose. Discussion & contributions welcome.

See the [LICENSE](LICENSE) file for licensing terms.

## Examples

See the [examples](examples) directory for a very small demo:

- Input proto definition in [example.proto](examples/example.proto)
- Different releases of the example in the [examples/releases](examples/releases) directory
- Generated code and a server in the [examples/server](examples/server) directory

A more complete suite of examples can be found in the [fixtures](fixtures)
directory.

## In short

### Proto releases

A message can be annotated with release gates as follows:

```proto
message Hello {
    int32 unreleased = 1;
    string released = 2 [
        (releases.field).release_in = 12;
    ];
    float32 preview_only = 3 [
        (releases.field).preview_in = 17;
    ];
}
```

Releases can be generated specifying the number of a release and whether it is a
preview release or not. In above example, only releases from #12 onwards will
include the `released` field, preview-releases from #17 onwards will include the
`preview_only` field, and release #0 will include all fields. In all cases, the
release options are stripped.

Use something like the following commands to generate the releases (assuming
protoc is already installed):

```sh
go install github.com/devnev/proto-releases/protoc-gen-release
protoc my_api.proto --release_out=releases/alpha
protoc my_api.proto --release_out=releases/beta17 --release_opt=release=17,preview=true
protoc my_api.proto --release_out=releases/stable17 --release_opt=release=17
```

`releases/alpha/my_api.proto`:

```proto
message Hello {
    int32 unreleased = 1;
    string released = 2;
    float32 preview_only = 3;
}
```

`releases/beta17/my_api.proto`:

```proto
message Hello {
    int32 unreleased = 1;
    string released = 2;
    float32 preview_only = 3;
}
```

`releases/stable17/my_api.proto`:

```proto
message Hello {
    int32 unreleased = 1;
    string released = 2;
}
```

### Automatic release implementation

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
    beta17.RegisterApiServer(srv, beta17.BaseApiServer{
        Base: apiSrv
    })
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
