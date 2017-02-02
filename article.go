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
	PublishedAt     string   `json:"publishedAt"`
	License         string   `json:"license"`
	NotifyFollowers bool     `json:"notifyFollowers"`
}

// PostedArticle represents an article on
type PostedArticle struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	AuthorID      string   `json:"authorId"`
	Tags          []string `json:"tags"`
	URL           string   `json:"url"`
	CanonicalURL  string   `json:"canonicalUrl"`
	PublishStatus string   `json:"publishStatus"`
	PublishAt     int      `json:"publishedAt"`
	License       string   `json:"license"`
	LicenseURL    string   `json:"licenseUrl"`
}
