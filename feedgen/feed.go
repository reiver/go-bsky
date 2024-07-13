package feedgen

import (
	"encoding/json"
)

type Feed interface {
	json.Marshaler
	FeedLen() int
	Items() []Item
}

type internalFeed struct {
	items []Item
}

var _ json.Marshaler = internalFeed{}

func FeedItems(items ...Item) Feed {
	return internalFeed{
		items:items,
	}
}

func (receiver internalFeed) FeedLen() int {
	return len(receiver.items)
}

func (receiver internalFeed) Items() []Item {
	return receiver.items
}

func (receiver internalFeed) MarshalJSON() ([]byte, error) {
	var data []Item = receiver.Items()

	if len(data) <= 0 {
		return []byte("[]"), nil
	}

	return json.Marshal(data)
}
