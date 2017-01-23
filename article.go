// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

type Article struct {
	ID            string   `json:""id"`
	Title         string   `json:""title"`
	AuthorID      string   `json:""authorId"`
	Tags          []string `json:""tags"`
	URL           string   `json:""url"`
	CanonicalURL  string   `json:""canonicalUrl"`
	PublishStatus string   `json:""publishStatus"`
	PublishAt     int      `json:""publishedAt"`
	License       string   `json:""license"`
	LicenseURL    string   `json:""licenseUrl"`
}
