package feedgen_test

import (
	"testing"

	"bytes"
	"encoding/json"

	"github.com/reiver/go-bsky/feedgen"
)

func TestReason_MarshalJSON(t *testing.T) {

	tests := []struct{
		Reason feedgen.Reason
		Expected []byte
	}{
		{
			Reason: feedgen.ReasonRepost("at://here"),
			Expected: []byte(`{"$type":"app.bsky.feed.defs#skeletonReasonRepost","repost":"at://here"}`),
		},
		{
			Reason: feedgen.ReasonRepost("at://once:twice:thrice:fource"),
			Expected: []byte(`{"$type":"app.bsky.feed.defs#skeletonReasonRepost","repost":"at://once:twice:thrice:fource"}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.Reason)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("ITEM: %#v", test.Reason)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d,the actual json-marshaled item is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				continue
			}
		}
	}
}
