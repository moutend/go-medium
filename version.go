// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

var (
	version  = "v0.3.0"
	revision = "latest"
)

// Version returns version and revision information.
func Version() string {
	return version + "-" + revision
}
