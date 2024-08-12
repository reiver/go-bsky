package firehose

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilDecoder       = erorr.Error("bsky: nil decoder")
	errNilMessageHeader = erorr.Error("bsky: nil message-header")
	errNilPayload       = erorr.Error("bsky: nil payload")
	errNilReader        = erorr.Error("bsky: nil reader")
	errNilReceiver      = erorr.Error("bsky: nil receiver")
	errNoBlocks         = erorr.Error("bsky: no blocks")
	errBlocksNotBytes   = erorr.Error("bsky: blocks not bytes")
)
