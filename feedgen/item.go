package feedgen

import (
	"encoding/json"
)

type Item interface {
	json.Marshaler
	ItemAttributes() map[string]interface{}
}
