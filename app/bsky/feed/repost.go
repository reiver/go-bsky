package feed

import (
//	goerrors "errors"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-bsky/internal/errors"
//	"github.com/reiver/go-bsky/internal/strng"
)

const RepostTypeValue = "app.bsky.feed.repost"

// Repost represents a 'app.bsky.feed.repost'.
type Repost struct {
	Type      RepostType `json:"$type"`
	CreatedAt string     `json:"createdAt"`
	Subject   struct{
		URI string   `json:"uri"`
		CID string   `json:"cid"`
	}                    `json:"subject"`
}

type RepostType struct {}

func (RepostType) String() string {
	return RepostTypeValue
}

func (RepostType) MarshalText() ([]byte, error) {
	return []byte(RepostTypeValue), nil
}

func (receiver *RepostType) UnmarshalText(text []byte) error {
	const expectedType = RepostTypeValue

	if nil == receiver {
		return errors.ErrNilReceiver
	}

	var actualType string = string(text)

	if expectedType != actualType {
		return erorr.Errorf("bsky: only the value %q can be text-unmarshaled into something of type %T", expectedType, *receiver)
	}

	return nil
}
