package record

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilSource = erorr.Error("bsky: nil source")
	errNotTyped  = erorr.Error("bsky: not found")
)
