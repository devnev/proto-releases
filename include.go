// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package releases

func Include(range_ *Range, config *Config) bool {
	if config.GetRelease() <= 0 {
		// no current release specified so include everything
		return true
	}
	if range_.GetRemoveIn() > 0 && range_.GetRemoveIn() <= config.GetRelease() {
		// skip if removed
		return false
	}
	if range_.GetReleaseIn() > 0 && range_.GetReleaseIn() <= config.GetRelease() {
		// keep if released
		return true
	}
	if config.GetPreview() && range_.GetPreviewIn() > 0 && range_.GetPreviewIn() <= config.GetRelease() {
		// keep if previewed
		return true
	}
	return false
}
