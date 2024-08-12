package firehose

// The message that comes back from the Bluesky Firehose websocket is 2 CBOR objects concatenated with each other.
// The first part is called the message-header. The second part is called the message-payload.
//
// MessagePayload represents the message-payload.
type MessagePayload map[string]any

func (receiver MessagePayload) Rebase() (bool, bool) {
	const name string = "rebase"
	var empty bool

	if nil == receiver {
		return empty, false
	}

	value, found := receiver["rebase"]
	if !found {
		return empty, false
	}

	switch casted := value.(type) {
	case bool:
		return casted, true
	default:
		return empty,  false
	}
}

func (receiver MessagePayload) Rev() (string, bool) {
	const name string = "rev"
	var empty string

	if nil == receiver {
		return empty, false
	}

	value, found := receiver[name]
	if !found {
		return empty, false
	}

	switch casted := value.(type) {
	case string:
		return casted, true
	default:
		return empty,  false
	}
}

func (receiver MessagePayload) Seq() (int, bool) {
	const name string = "seq"
	var empty int

	if nil == receiver {
		return empty, false
	}

	value, found := receiver[name]
	if !found {
		return empty, false
	}

	switch casted := value.(type) {
	case int:
		return casted, true
	default:
		return empty,  false
	}
}

func (receiver MessagePayload) Since() (string, bool) {
	const name string = "since"
	var empty string

	if nil == receiver {
		return empty, false
	}

	value, found := receiver[name]
	if !found {
		return empty, false
	}

	switch casted := value.(type) {
	case string:
		return casted, true
	default:
		return empty,  false
	}
}

func (receiver MessagePayload) TooBig() (bool, bool) {
	const name string = "tooBig"
	var empty bool

	if nil == receiver {
		return empty, false
	}

	value, found := receiver[name]
	if !found {
		return empty, false
	}

	switch casted := value.(type) {
	case bool:
		return casted, true
	default:
		return empty,  false
	}
}

func (receiver MessagePayload) Time() (string, bool) {
	const name string = "time"
	var empty string

	if nil == receiver {
		return empty, false
	}

	value, found := receiver[name]
	if !found {
		return empty, false
	}

	switch casted := value.(type) {
	case string:
		return casted, true
	default:
		return empty,  false
	}
}
