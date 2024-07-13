package feedgen

import (
	"encoding/json"

	"github.com/reiver/go-opt"
)

type FeedSkeleton interface {
	json.Marshaler
	Cursor() (string,bool)
	Feed() Feed
}

type internalFeedSkeleton struct {
	cursor opt.Optional[string]
	feed   Feed
}

var _ json.Marshaler = internalFeedSkeleton{}

func FeedSkeletonItems(items ...Item) FeedSkeleton {
	return internalFeedSkeleton {
		feed: FeedItems(items...),
	}
}

func FeedSkeletonCursonItems(cursor string, items ...Item) FeedSkeleton {
	return internalFeedSkeleton {
		cursor: opt.Something(cursor),
		feed: FeedItems(items...),
	}
}

func (receiver internalFeedSkeleton) Cursor() (string, bool) {
	return receiver.cursor.Get()
}

func (receiver internalFeedSkeleton) Feed() Feed {
	return receiver.feed
}

func (receiver internalFeedSkeleton) MarshalJSON() ([]byte, error) {
	var data = map[string]interface{}{}

	data["feed"] = receiver.Feed()

	{
		cursor, something := receiver.Cursor()
		if something {
			data["cursor"] = cursor
		}
	}

	return json.Marshal(data)
}
