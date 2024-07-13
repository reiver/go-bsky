package feedgen_test

import (
	"testing"

	"bytes"
	"encoding/json"

	"github.com/reiver/go-bsky/feedgen"
)

func TestFeed_MarshalJSON(t *testing.T) {

	tests := []struct{
		Feed feedgen.Feed
		Expected []byte
	}{
		{
			Feed: nil,
			Expected: []byte(`null`),
		},
		{
			Feed: feedgen.FeedItems(),
			Expected: []byte(`[]`),
		},



		{
			Feed: feedgen.FeedItems(
				feedgen.ItemPost("at://apple:banana:cherry"),
			),
			Expected: []byte(
				`[`+
					`{"post":"at://apple:banana:cherry"}`+
				`]`,
			),
		},
		{
			Feed: feedgen.FeedItems(
				feedgen.ItemPostReason("at://apple:banana:cherry", feedgen.ReasonRepost("at://here")),
			),
			Expected: []byte(
				`[`+
					`{`+
						`"post":"at://apple:banana:cherry"`+
						`,`+
						`"reason":{"$type":"app.bsky.feed.defs#skeletonReasonRepost","repost":"at://here"}`+
					`}`+
				`]`,
			),
		},
		{
			Feed: feedgen.FeedItems(
				feedgen.ItemPostReason("at://apple:banana:cherry", feedgen.ReasonRepost("at://once:twice:thrice:fource")),
			),
			Expected: []byte(
				`[`+
					`{`+
						`"post":"at://apple:banana:cherry"`+
						`,`+
						`"reason":{"$type":"app.bsky.feed.defs#skeletonReasonRepost","repost":"at://once:twice:thrice:fource"}`+
					`}`+
				`]`,
			),
		},



		{
			Feed: feedgen.FeedItems(
				feedgen.ItemPost("at://once:twice:thrice:fource"),
				feedgen.ItemPostReason("at://apple:banana:cherry", feedgen.ReasonRepost("at://here")),
			),
			Expected: []byte(
				`[`+
					`{"post":"at://once:twice:thrice:fource"}`+
					`,`+
					`{`+
						`"post":"at://apple:banana:cherry"`+
						`,`+
						`"reason":{"$type":"app.bsky.feed.defs#skeletonReasonRepost","repost":"at://here"}`+
					`}`+
				`]`,
			),
		},
		{
			Feed: feedgen.FeedItems(
				feedgen.ItemPost("at://once:twice:thrice:fource"),
				feedgen.ItemPostReason("at://apple:banana:cherry", feedgen.ReasonRepost("at://here")),
				feedgen.ItemPostReason("at://one:two:three:four", feedgen.ReasonRepost("at://hello.world")),
			),
			Expected: []byte(
				`[`+
					`{"post":"at://once:twice:thrice:fource"}`+
					`,`+
					`{`+
						`"post":"at://apple:banana:cherry"`+
						`,`+
						`"reason":{"$type":"app.bsky.feed.defs#skeletonReasonRepost","repost":"at://here"}`+
					`}`+
					`,`+
					`{`+
						`"post":"at://one:two:three:four"`+
						`,`+
						`"reason":{"$type":"app.bsky.feed.defs#skeletonReasonRepost","repost":"at://hello.world"}`+
					`}`+
				`]`,
			),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.Feed)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("ITEM: %#v", test.Feed)
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
