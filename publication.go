// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

// Publication represents a Medium Publication.
type Publication struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	ImageURL    string `json:"imageUrl"`
	client      *Client
}
