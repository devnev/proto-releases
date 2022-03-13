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

A more complete suite of examples can be found in the [fixtures](fixtures) directory.

## In short

### Proto releases

A message can be annotated with release gates as follows:

```proto
message Hello {
    int32 unreleased = 1;
    string released = 2 [
        (releases.field).release_in = 14;
    ];
    float32 preview_only = 3 [
        (releases.field).preview_in = 22;
    ];
}
```

Releases can be generated specifying the number of a release and whether it is a preview release or not. In above example, only releases from #22 onwards will include the `released` field, preview-releases from #14 onwards will include the `preview_only` field, and release 0 will include all fields. In all cases, the release options are stripped.

Use something like the following commands to generate the releases (assuming protoc is already installed):

```sh
go install github.com/devnev/proto-releases/protoc-gen-release
protoc my_api.proto --release_out=releases/alpha
protoc my_api.proto --release_out=releases/beta17 --release_opt=release=17,preview=true
protoc my_api.proto --release_out=releases/stable17 --release_opt=release=17,preview=true
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
go install \
    github.com/devnev/proto-releases/protoc-gen-release

protoc my_api.proto \
    --release_out=releases/beta17 \
    --release_opt=release=17,preview=true,go_package=github.com/me/releases/beta17

protoc my_api.proto \
    --go_out=server --go_opt=paths=source_relative \
    --go-grpc_out=server --go-grpc_opt=paths=source_relative

go install \
    github.com/devnev/proto-releases/protoc-gen-go-baseconvert

protoc beta17/my_api.proto \
    -Ireleases \
    --go_out=server --go_opt=paths=source_relative \
    --go-grpc_out=server --go-grpc_opt=paths=source_relative \
    --go-baseconvert_out=server \
    --go-baseconvert_opt=base=github.com/me/server,paths=source_relative
```
