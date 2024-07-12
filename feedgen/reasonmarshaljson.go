package feedgen

import (
	"encoding/json"
)

func reasonMarshalJSON(reason Reason) ([]byte, error) {
	if nil == reason {
		return nil, errNilReason
	}

	var data map[string]interface{} = reason.ReasonAttributes()
	if nil == data {
		data = map[string]interface{}{}
	}

	data["$type"] = reason.ReasonType()

	return json.Marshal(data)
}
