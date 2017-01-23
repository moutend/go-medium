// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

package medium

import (
	"encoding/json"
	"fmt"
	"io"
)

func decodeJSON(body io.Reader, out interface{}) (err error) {
	decoder := json.NewDecoder(body)
	err = decoder.Decode(out)
	if err != nil {
		err = fmt.Errorf("failed to decode JSON.\n%s\n", err)
	}
	return
}
