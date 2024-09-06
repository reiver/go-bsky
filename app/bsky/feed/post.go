package feed

import (
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-bsky/internal/errors"
)

const PostTypeValue = "app.bsky.feed.post"

// Post represents a 'app.bsky.feed.post'.
type Post struct {
	Type      PostType `json:"$type"`
	CreatedAt string   `json:"createdAt"`
	Text      string   `json:"text"`
}

type PostType struct {}

func (PostType) String() string {
	return PostTypeValue
}

func (PostType) MarshalText() ([]byte, error) {
	return []byte(PostTypeValue), nil
}

func (receiver *PostType) UnmarshalText(text []byte) error {
	const expectedType = PostTypeValue

	if nil == receiver {
		return errors.ErrNilReceiver
	}

	var actualType string = string(text)

	if expectedType != actualType {
		return erorr.Errorf("bsky: only the value %q can be text-unmarshaled into something of type %T", expectedType, *receiver)
	}

	return nil
}
