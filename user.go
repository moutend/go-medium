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
	client   *Client
}

// Post posts an article to the authenticated user's profile.
func (u *User) Post(a Article) (*PostedArticle, error) {
	return u.client.postArticle("users", u.ID, a)
}

// Publications returns specified user's publications.
func (u *User) Publications() (p []*Publication, err error) {
	req, err := u.client.newRequest("GET", "/users/"+u.ID+"/publications", nil)
	if err != nil {
		return
	}
	r, err := u.client.do(req)
	if err != nil {
		return
	}
	return r.Publications(u.client)
}
