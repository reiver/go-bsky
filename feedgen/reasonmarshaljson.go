package feedgen

import (
	"encoding/json"
)

func reasonMarshalJSON(reason Reason) ([]byte, error) {
	if nil == reason {
		return nil, errNilReason
	}

	var data map[string]interface{} = reason.ReasonAttributes()

	data["$type"] = reason.ReasonType()

	return json.Marshal(data)
}
