package feed

import (
	"encoding/json"

	"testing"
)

func TestTest(t *testing.T) {

	var post Post

	p, err := json.Marshal(post)
	if nil != err {
		t.Errorf("ERROR: (%T) %s", err, err)
		return
	}

	t.Errorf("JSON:\n%s", string(p))
}
