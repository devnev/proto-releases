// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package releases

//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative releases.proto config.proto
