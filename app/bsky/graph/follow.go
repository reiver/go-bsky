package graph

import (
	goerrors "errors"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-bsky/internal/errors"
	"github.com/reiver/go-bsky/internal/strng"
	"github.com/reiver/go-bsky/record/registry"
)

const FollowTypeValue = "app.bsky.graph.follow"

func init() {
	registry.NewFuncs.Set(FollowTypeValue, func()registry.Record {
		return new(Follow)
	})
}

// Follow represents a 'app.bsky.graph.follow'.
type Follow struct {
	Type      FollowType `json:"$type"`
	CreatedAt string     `json:"createdAt"`
	Subject   string     `json:"subject"`
}

type FollowType struct {}

func (FollowType) String() string {
	return FollowTypeValue
}

func (FollowType) MarshalText() ([]byte, error) {
	return []byte(FollowTypeValue), nil
}

func (receiver *FollowType) UnmarshalText(text []byte) error {
	const expectedType = FollowTypeValue

	if nil == receiver {
		return errors.ErrNilReceiver
	}

	var actualType string = string(text)

	if expectedType != actualType {
		return erorr.Errorf("bsky: only the value %q can be text-unmarshaled into something of type %T", expectedType, *receiver)
	}

	return nil
}

func (receiver *Follow) FromMap(src map[string]any) error {
	if nil == receiver {
		return errors.ErrNilReceiver
	}

	err1 := strng.Set(&(receiver.Type),      src["$type"])
	err2 := strng.Set(&(receiver.CreatedAt), src["createdAt"])
	err3 := strng.Set(&(receiver.Subject),   src["subject"])

	var err error = goerrors.Join(err1, err2, err3)

	return err
}

func (receiver Follow) RecordType() string {
	return receiver.Type.String()
}
