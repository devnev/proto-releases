# Updating HTTP API Annotations

A method that is included in a release and has a `(google.api.http)` annotation can be rewritten to have its path prefixed:

```proto
service ExampleService {
    rpc SayHello(google.protobuf.Empty) returns (google.protobuf.Empty) {
        option (releases.method).release_in = 2;
        option (google.api.http) = {
            get: "/say-hello"
        };
    }
}
```

Released with

```sh
protoc \
    my_api.proto \
    --release_out=releases/stable \
    --release_opt=release=2 \
    --release_opt=http_rule=:/stable
```

Will produce

```proto
service ExampleService {
    rpc SayHello(google.protobuf.Empty) returns (google.protobuf.Empty) {
        option (google.api.http) = {
            get: "/stable/say-hello"
        };
    }
}
```
