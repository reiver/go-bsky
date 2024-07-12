package feedgen_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-bsky/feedgen"
)

func TestReason_ReasonAttributes(t *testing.T) {

	tests := []struct{
		Reason feedgen.Reason
		Expected map[string]interface{}
	}{
		{
			Reason: feedgen.ReasonRepost("at://here"),
			Expected: map[string]interface{}{
				"repost":            "at://here",
			},
		},
	}

	for testNumber, test := range tests {

		actual := test.Reason.ReasonAttributes()

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d,the actual reason-type is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
