// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

import (
	"strings"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	json := `{"message": "foobar"}`

	var e struct {
		Message string
	}
	err := decodeJSON(strings.NewReader(json), &e)
	if err != nil {
		t.Fatal(err)
	}
	expected := "foobar"
	actual := e.Message
	if expected != actual {
		t.Fatalf("\nexpected: %s\nactual: %s\n", expected, actual)
	}
	return
}
