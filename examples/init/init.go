package init

import "os"

func init() {
	// Not mentioned by documentation, the actual proto file paths also have to
	// be unique, at least in the Go implementation. Lower conflict detection to
	// warn for now while figuring out a solution.
	os.Setenv("GOLANG_PROTOBUF_REGISTRATION_CONFLICT", "warn")
}
