// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

import (
	"bytes"
	"encoding/json"
	"net/url"
)

// Publication represents a Medium Publication.
type Publication struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	ImageURL    string `json:"imageUrl"`
	client      *Client
}

// Post posts an article to the authenticated user's publication.
func (p *Publication) Post(a Article) (pa *PostedArticle, err error) {
	path, _ := url.Parse("/publications/" + p.ID + "/posts")
	content, err := json.Marshal(a)
	if err != nil {
		return
	}
	r, err := p.client.post(path, bytes.NewReader(content))
	return r.PostedArticle(p.client)
}
