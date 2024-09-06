package record


import (
	"github.com/reiver/go-erorr"
)

func Type(src any) (string, error) {
	const key string = "$type"

	var empty string

	if nil == src {
		return empty, errNilSource
	}

	var typ any

	switch casted := src.(type) {
	case map[string]string:
		var found bool

		typ, found = casted[key]
		if !found {
			return empty, errNotTyped
		}
	case map[string]any:
		var found bool

		typ, found = casted[key]
		if !found {
			return empty, errNotTyped
		}
	case map[any]string:
		var found bool

		typ, found = casted[key]
		if !found {
			return empty, errNotTyped
		}
	case map[any]any:
		var found bool

		typ, found = casted[key]
		if !found {
			return empty, errNotTyped
		}
	default:
		return empty, erorr.Errorf("bsky: cannot get value of \"$type\" field for something of type %T", src)
	}

	{
		switch casted := typ.(type) {
		case string:
			return casted, nil
		default:
			return empty, erorr.Errorf("bsky: expected \"$type\" field to be of type %T but actually is of type %T", empty, typ)
		}
	}
}

