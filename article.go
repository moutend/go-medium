// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

// Article represents an article.
type Article struct {
	Title           string   `json:"title"`
	ContentFormat   string   `json:"contentFormat"`
	Content         string   `json:"content"`
	CanonicalURL    string   `json:"canonicalUrl"`
	Tags            []string `json:"tags"`
	PublishStatus   string   `json:"publishStatus"`
	License         string   `json:"license"`
	NotifyFollowers bool     `json:"notifyFollowers"`
}
