package feedgen_test

import (
	"testing"

	"bytes"
	"encoding/json"

	"github.com/reiver/go-bsky/feedgen"
)

func TestItem_MarshalJSON(t *testing.T) {

	tests := []struct{
		Item feedgen.Item
		Expected []byte
	}{
		{
			Item: feedgen.ItemPost("at://apple:banana:cherry"),
			Expected: []byte(`{"post":"at://apple:banana:cherry"}`),
		},
		{
			Item: feedgen.ItemPostReason("at://apple:banana:cherry", feedgen.ReasonRepost("at://here")),
			Expected: []byte(`{"post":"at://apple:banana:cherry","reason":{"$type":"app.bsky.feed.defs#skeletonReasonRepost","repost":"at://here"}}`),
		},
		{
			Item: feedgen.ItemPostReason("at://apple:banana:cherry", feedgen.ReasonRepost("at://once:twice:thrice:fource")),
			Expected: []byte(`{"post":"at://apple:banana:cherry","reason":{"$type":"app.bsky.feed.defs#skeletonReasonRepost","repost":"at://once:twice:thrice:fource"}}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.Item)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("ITEM: %#v", test.Item)
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
