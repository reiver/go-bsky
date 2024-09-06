package feed_test

import (
	"testing"

	app_bsky_feed "github.com/reiver/go-bsky/app/bsky/feed"
)

func TestRepostTypeValue_String(t *testing.T) {

	var specialType app_bsky_feed.RepostType

	actual := specialType.String()

	expected := app_bsky_feed.RepostTypeValue

	if expected != actual {
		t.Errorf("The actual .String() result for %T is not what was expected.", specialType)
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}

func TestRepostTypeValue_MarshalText(t *testing.T) {

	var specialType app_bsky_feed.RepostType

	actualBytes, err := specialType.MarshalText()
	if nil != err {
		t.Error("Did not expect to get an error but actually got one.")
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}

	actual := string(actualBytes)

	expected := app_bsky_feed.RepostTypeValue

	if expected != actual {
		t.Errorf("The actual .MarshalText() result for %T is not what was expected.", specialType)
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}

func TestRepostTypeValue_UnarshalText_success(t *testing.T) {

	var src []byte = []byte(app_bsky_feed.RepostTypeValue)

	var specialType app_bsky_feed.RepostType

	err := specialType.UnmarshalText(src)
	if nil != err {
		t.Error("Did not expect to get an error but actually got one.")
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}
}

func TestRepostTypeValue_UnarshalText_fail(t *testing.T) {

	var src []byte = []byte("THIS SHOULD FAIL!")

	var specialType app_bsky_feed.RepostType

	err := specialType.UnmarshalText(src)
	if nil == err {
		t.Error("Expected an error but did not actually get one.")
		return
	}

	{
		actual := err.Error()
		expected := `bsky: only the value "app.bsky.feed.repost" can be text-unmarshaled into something of type feed.RepostType`

		if expected != actual {
			t.Errorf("The actual error is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}
}
