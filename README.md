# Proto-Releases

Proto-releases provides protobuf extensions and tools to manage [stability
levels](https://google.aip.dev/181) and releases of protobuf-based APIs.

> This is experimental, work in progress, and by no means a complete set of tools
> for the above purpose. Discussion & contributions welcome.
>
> See the [LICENSE](LICENSE) file for licensing terms.

For example, this "internal" proto `acmecorp/hello.proto`:

```proto
package acmecorp;
message Hello {
    int32 repeat = 1 [
        (releases.field).preview_in = 9;
    ];
    string name = 2 [
        (releases.field).release_in = 5;
    ];
}
service GreetingService {
    rpc SayHello(Hello) returns google.protobuf.Empty {
        option (releases.method).release_in = 5;
        option (google.api.http) = {
            get: "/v1unstable/hello"
        };
    };
}
```

Can be converted to the following "stable" proto `acmecorp/v1/hello.proto`:

```proto
package acmecorp.v1;
message Hello {
    string name = 2;
}
service GreetingService {
    rpc SayHello(Hello) returns google.protobuf.Empty {
        option (google.api.http) = {
            get: "/v1/hello"
        };
    };
}
```

Further, a single implementation of the original service can be used in Go to
implement all released service versions:

```go
srv := grpc.NewServer()
greetingsSvc := NewGreetingsService()
RegisterGreetingsServiceServer(srv, greetingsSvc)
beta.RegisterGreetingsServiceBaseServer(srv, greetingsSvc)
stable.RegisterGreetingsServiceBaseServer(srv, greetingsSvc)
```

## Working Example

See the [examples](examples) directory for a very small demo:

- Input proto definition in [example.proto](examples/example.proto)
- Different releases of the example in the [examples/releases](examples/releases) directory
- Generated code and a server in the [examples/server](examples/server) directory

A more complete suite of examples can be found in the [fixtures](fixtures)
directory.

## Further Documentation

- [Generating a release](docs/generating_a_release.md)
- [Automatic release implementation](docs/automatic_release_implementation.md)
- [Updating HTTP API Annotations](docs/updating_http_api_annotations.md)
