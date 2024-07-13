package feedgen

import (
	"encoding/json"

	"github.com/reiver/go-opt"
)

type internalItem struct {
	post   opt.Optional[string]
	reason Reason
}

var _ json.Marshaler = internalItem{}

func ItemPost(uri string) Item {
	return internalItem{
		post: opt.Something(uri),
	}
}

func ItemPostReason(uri string, reason Reason) Item {
	return internalItem{
		post: opt.Something(uri),
		reason: reason,
	}
}

func (receiver internalItem) ItemAttributes() map[string]interface{} {

	var data = map[string]interface{}{}

	{
		post, something := receiver.post.Get()

		if !something {
			return nil
		}

		data["post"] = post
	}

	{
		reason := receiver.reason

		if nil != reason {
			data["reason"] = reason
		}
	}

	return data
}

func (receiver internalItem) MarshalJSON() ([]byte, error) {
	return itemMarshalJSON(receiver)
}
