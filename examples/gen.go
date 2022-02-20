// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package example

//go:generate go run github.com/devnev/proto-releases/gen-release -inc .. -out alpha example.proto
//go:generate go run github.com/devnev/proto-releases/gen-release -inc .. -out beta1 -rel 1 -pre example.proto
//go:generate go run github.com/devnev/proto-releases/gen-release -inc .. -out beta2 -rel 2 -pre example.proto
//go:generate go run github.com/devnev/proto-releases/gen-release -inc .. -out beta3 -rel 3 -pre example.proto
//go:generate go run github.com/devnev/proto-releases/gen-release -inc .. -out stable1 -rel 1 example.proto
//go:generate go run github.com/devnev/proto-releases/gen-release -inc .. -out stable2 -rel 2 example.proto
//go:generate go run github.com/devnev/proto-releases/gen-release -inc .. -out stable3 -rel 3 example.proto
