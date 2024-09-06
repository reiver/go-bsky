package feed

import (
	goerrors "errors"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-fallback"
	"github.com/reiver/go-maps"

	"github.com/reiver/go-bsky/internal/errors"
	"github.com/reiver/go-bsky/internal/strng"
	"github.com/reiver/go-bsky/record/registry"
)

const LikeTypeValue = "app.bsky.feed.like"

func init() {
	registry.NewFuncs.Set(LikeTypeValue, func()registry.Record {
		return new(Like)
	})
}

// Like represents a 'app.bsky.feed.like'.
type Like struct {
	Type      LikeType `json:"$type"`
	CreatedAt string   `json:"createdAt"`
	Subject   struct{
		CID string `json:"cid"`
		URI string `json:"uri"`
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

func (receiver *Like) FromMap(src map[string]any) error {
	if nil == receiver {
		return errors.ErrNilReceiver
	}

	err1 := strng.Set(&(receiver.Type),        src["$type"])
	err2 := strng.Set(&(receiver.CreatedAt),   src["createdAt"])
	err3 := strng.Set(&(receiver.Subject.CID), fallback.DefaultTo[string]{""}.WrapBool(maps.PathQueryCast[string](src, "subject", "cid")) )
	err4 := strng.Set(&(receiver.Subject.URI), fallback.DefaultTo[string]{""}.WrapBool(maps.PathQueryCast[string](src, "subject", "uri")) )

	var err error = goerrors.Join(err1, err2, err3, err4)

	return err
}

func (receiver Like) RecordType() string {
	return receiver.Type.String()
}
