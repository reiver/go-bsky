package feedgen_test

import (
	"testing"

	"github.com/reiver/go-bsky/feedgen"
)

func TestReason_ReasonType(t *testing.T) {

	tests := []struct{
		Reason feedgen.Reason
		Expected string
	}{
		{
			Reason: feedgen.ReasonRepost("at://here"),
			Expected: "app.bsky.feed.defs#skeletonReasonRepost",
		},
	}

	for testNumber, test := range tests {

		actual := test.Reason.ReasonType()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d,the actual reason-type is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
