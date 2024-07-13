package feedgen

import (
	"encoding/json"
)

func itemMarshalJSON(reason Item) ([]byte, error) {
	if nil == reason {
		return nil, errNilReason
	}

	var data map[string]interface{} = reason.ItemAttributes()
	if nil == data {
		data = map[string]interface{}{}
	}

	return json.Marshal(data)
}
