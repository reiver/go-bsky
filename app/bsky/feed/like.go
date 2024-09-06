package feed

import (
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-bsky/internal/errors"
)

const LikeTypeValue = "app.bsky.feed.like"

// Like represents a 'app.bsky.feed.like'.
type Like struct {
	Type      LikeType `json:"$type"`
	CreatedAt string   `json:"createdAt"`
	Subject   struct{
		URI string `json:"uri"`
		CID string `json:"cid"`
	}                  `json:"subject"`
}

type LikeType struct {}

func (LikeType) String() string {
	return LikeTypeValue
}

func (LikeType) MarshalText() ([]byte, error) {
	return []byte(LikeTypeValue), nil
}

func (receiver *LikeType) UnmarshalText(text []byte) error {
        const expectedType = LikeTypeValue

        if nil == receiver {
                return errors.ErrNilReceiver
        }

        var actualType string = string(text)

        if expectedType != actualType {
		return erorr.Errorf("bsky: only the value %q can be text-unmarshaled into something of type %T", expectedType, *receiver)
         }

        return nil
}
