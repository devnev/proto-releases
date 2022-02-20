# Proto-Releases

Proto-releases aims to provide protobuf extensions and tools to manage releases
of proto definitions. It currently allows creating unstable/alpha, preview/beta
and stable release channels of the same proto definitions.

This is experimental, work in progress, and by no means a complete set of tools
for the above purpose. Discussion & contributions welcome.

See the [LICENSE](LICENSE) file for licensing terms.

## Unstable/Alpha, Preview/Beta, Stable/Released

Within a major release, protocol changes may be promoted from unstable (e.g. for
internal use and testing or early adopters), to a preview audience (e.g. beta
program), to stable & generally available.

This is modeled using extensions that indicate in which release a protobuf
element is intended to be available for preview or general availability. The
gen-release tool can be used to generate derived proto definitions for each
case.

See the [examples](examples) directory for an input proto definition
([example.proto](examples/example.proto)), commands (in
[gen.go](examples/gen.go)) and resulting releases (in corresponding
subdirectories).
