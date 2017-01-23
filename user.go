// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

import (
	"bytes"
	"encoding/json"
	"net/url"
)

// User represents a Medium user
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"imageUrl"`
	client   *Client
}

// Post posts an article to the authenticated user's profile.
func (u *User) Post(a Article) (*PostedArticle, error) {
	path, _ := url.Parse("/users/" + u.ID + "/posts")
	content, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	r, err := u.client.post(path, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	return r.PostedArticle(u.client)
}

// Publications returns specified user's publications.
func (u *User) Publications() (p []*Publication, err error) {
	path, _ := url.Parse("/users/" + u.ID + "/publications")
	r, err := u.client.get(path)
	if err != nil {
		return
	}
	return r.Publications(u.client)
}
