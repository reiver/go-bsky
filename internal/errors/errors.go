package errors

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrNilDestination = erorr.Error("bsky: nil destination")
	ErrNilReceiver    = erorr.Error("bsky: nil receiver")
	ErrNilSource      = erorr.Error("bsky: nil source")
)
