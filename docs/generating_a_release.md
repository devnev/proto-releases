# Generating a Release

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
