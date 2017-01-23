// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

import (
	"bytes"
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

func (r rawbody) Article(c *Client) (*Article, error) {
	var i struct {
		Data *Article
	}
	err := decodeJSON(bytes.NewReader(r), &i)
	return i.Data, err
}

func (r rawbody) Publications(c *Client) ([]*Publication, error) {
	var i struct {
		Data []*Publication
	}
	err := decodeJSON(bytes.NewReader(r), &i)
	for n, _ := range i.Data {
		i.Data[n].client = c
	}
	return i.Data, err
}

// Post creates a post on the authenticated user's profile page.
func (u *User) Post(post Post) (a *Article, err error) {
	path, _ := url.Parse("/users/" + u.ID + "/posts")
	r, err := u.client.post(path, post)
	return r.Article(u.client)
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
