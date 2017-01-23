// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

import (
	"bytes"
	"fmt"
)

type rawbody []byte

//debug: {"errors":[{"message":"Invalid publishStatus specified.","code":2008}]}
func (r rawbody) Error() error {
	var i struct {
		Errors []Error
	}
	err := decodeJSON(bytes.NewReader(r), &i)
	if err != nil {
		return err
	}
	if len(i.Errors) > 0 {
		return fmt.Errorf("%s (code:%d)", i.Errors[0].Message, i.Errors[0].Code)
	} else {
		return fmt.Errorf("broken response")
	}
}

func (r rawbody) User(c *Client) (*User, error) {
	var i struct {
		Data User
	}
	err := decodeJSON(bytes.NewReader(r), &i)
	i.Data.client = c
	return &i.Data, err
}

func (r rawbody) PostedArticle(c *Client) (*PostedArticle, error) {
	var i struct {
		Data *PostedArticle
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
