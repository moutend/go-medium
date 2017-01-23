// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

// Contributor represents a contributor for the specific publication.
type Contributor struct {
	PublicationID string `json:"publicationId"`
	UserID        string `json:"userId"`
	Role          string `json:"role"`
}
