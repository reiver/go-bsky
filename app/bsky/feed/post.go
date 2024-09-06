package feed

import (
	goerrors "errors"

	"github.com/reiver/go-atproto/record/registry"
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-bsky/internal/errors"
	"github.com/reiver/go-bsky/internal/strng"
)

const PostTypeValue = "app.bsky.feed.post"

func init() {
	registry.NewFuncs.Set(PostTypeValue, func()registry.Record {
		return new(Post)
	})
}

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

func (receiver *Post) FromMap(src map[string]any) error {
	if nil == receiver {
		return errors.ErrNilReceiver
	}

	err1 := strng.Set(&(receiver.Type),      src["$type"])
	err2 := strng.Set(&(receiver.CreatedAt), src["createdAt"])
	err3 := strng.Set(&(receiver.Text),      src["text"])

	var err error = goerrors.Join(err1, err2, err3)

	return err
}

func (receiver Post) RecordType() string {
	return receiver.Type.String()
}
