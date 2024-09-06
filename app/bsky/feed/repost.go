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

const RepostTypeValue = "app.bsky.feed.repost"

func init() {
	registry.NewFuncs.Set(RepostTypeValue, func()registry.Record {
		return new(Repost)
	})
}

// Repost represents a 'app.bsky.feed.repost'.
type Repost struct {
	Type      RepostType `json:"$type"`
	CreatedAt string     `json:"createdAt"`
	Subject   struct{
		CID string   `json:"cid"`
		URI string   `json:"uri"`
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

func (receiver *Repost) FromMap(src map[string]any) error {
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
