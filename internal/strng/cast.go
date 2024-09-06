package strng

import (
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-bsky/internal/errors"
)

func toBytes(src any) ([]byte, error) {
	var empty []byte

	if nil == src {
		return empty, errors.ErrNilSource
	}

	var value []byte

	switch casted := src.(type) {
	case string:
		value = []byte(casted)
	case []byte:
		value = casted
	case []rune:
		value = []byte(string(casted))
	default:
		return empty, erorr.Errorf("bsky: cannot (bytes) string-set from something of type %T", src)
	}

	return value, nil
}

func toRunes(src any) ([]rune, error) {
	var empty []rune

	if nil == src {
		return empty, errors.ErrNilSource
	}

	var value []rune

	switch casted := src.(type) {
	case string:
		value = []rune(casted)
	case []byte:
		value = []rune(string(casted))
	case []rune:
		value = casted
	default:
		return empty, erorr.Errorf("bsky: cannot (runes) string-set from something of type %T", src)
	}

	return value, nil
}

func toString(src any) (string, error) {
	var empty string

	if nil == src {
		return empty, errors.ErrNilSource
	}

	var value string

	switch casted := src.(type) {
	case string:
		value = casted
	case []byte:
		value = string(casted)
	case []rune:
		value = string(casted)
	default:
		return empty, erorr.Errorf("bsky: cannot (string) string-set from something of type %T", src)
	}

	return value, nil
}
