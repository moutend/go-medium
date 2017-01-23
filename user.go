// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

// User represents a Medium user
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"imageUrl"`
}
