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

// Post posts an article to the authenticated user's publication.
func (p *Publication) Post(a Article) (*PostedArticle, error) {
	return p.client.postArticle("publications", p.ID, a)
}

// Contributors returns a list of contributors for the publication.
func (p *Publication) Contributors() ([]*Contributor, error) {
	req, err := p.client.newRequest("GET", "/publications/"+p.ID+"/contributors", nil)
	if err != nil {
		return nil, err
	}
	r, err := p.client.do(req)
	if err != nil {
		return nil, err
	}
	return r.Contributors()
}
