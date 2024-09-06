package strng

import (
	"encoding"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-bsky/internal/errors"
)

func Set(dst any, src any) error {
	if nil == dst {
		return errors.ErrNilDestination
	}

	switch casted := dst.(type) {
	case encoding.TextUnmarshaler:
		return SetIntoTextUnmarshaler(casted, src)
	case *string:
		return SetIntoString(casted, src)
	case *[]byte:
		return SetIntoBytes(casted, src)
	case *[]rune:
		return SetIntoRunes(casted, src)
	default:
		return erorr.Errorf("bsky: cannot string-set into something of type %T", dst)
	}

	return nil
}

func SetIntoBytes(dst *[]byte, src any) error {
	if nil == dst {
		return errors.ErrNilDestination
	}

	value, err := toBytes(src)
	if nil != err {
		return nil
	}

	*dst = value
	return nil
}

func SetIntoRunes(dst *[]rune, src any) error {
	if nil == dst {
		return errors.ErrNilDestination
	}

	value, err := toRunes(src)
	if nil != err {
		return nil
	}

	*dst = value
	return nil
}

func SetIntoString(dst *string, src any) error {
	if nil == dst {
		return errors.ErrNilDestination
	}

	value, err := toString(src)
	if nil != err {
		return nil
	}

	*dst = value
	return nil
}

func SetIntoTextUnmarshaler(dst encoding.TextUnmarshaler, src any) error {
	if nil == dst {
		return errors.ErrNilDestination
	}

	value, err := toBytes(src)
	if nil != err {
		return nil
	}

	err = dst.UnmarshalText(value)
	if nil != err {
		return erorr.Errorf("bsky: cannot string-set into text-unmarshaler: %w", err)
	}
	return nil
}
