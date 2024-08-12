package firehose

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilDecoder       = erorr.Error("bsky: nil decoder")
	errNilMessageHeader = erorr.Error("bsky: nil message-header")
	errNilPayload       = erorr.Error("bsky: nil payload")
	errNilReader        = erorr.Error("bsky: nil reader")
)
