// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

var testdata = filepath.Join("testdata", "json")

func TestImage(t *testing.T) {
	var r rawbody
	r, err := ioutil.ReadFile(filepath.Join(testdata, "image.json"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Image()
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestContributors(t *testing.T) {
	var r rawbody
	r, err := ioutil.ReadFile(filepath.Join(testdata, "contributors.json"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Contributors()
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestError(t *testing.T) {
	var r rawbody
	r, err := ioutil.ReadFile(filepath.Join(testdata, "error.json"))
	if err != nil {
		t.Fatal(err)
	}
	err = r.Error()
	if err == nil {
		t.Fatalf("err should not be nil\n")
	}
	return
}

func TestUser(t *testing.T) {
	var r rawbody
	r, err := ioutil.ReadFile(filepath.Join(testdata, "user.json"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.User(nil)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestPostedArticle(t *testing.T) {
	var r rawbody
	r, err := ioutil.ReadFile(filepath.Join(testdata, "postedarticle.json"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.PostedArticle(nil)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestPublications(t *testing.T) {
	var r rawbody
	r, err := ioutil.ReadFile(filepath.Join(testdata, "publications.json"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Publications(nil)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestToken(t *testing.T) {
	var r rawbody
	r, err := ioutil.ReadFile(filepath.Join(testdata, "tokens.json"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Token()
	if err != nil {
		t.Fatal(err)
	}
	return
}
