package graph

import (
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-bsky/internal/errors"
)

const FollowTypeValue = "app.bsky.graph.follow"

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
