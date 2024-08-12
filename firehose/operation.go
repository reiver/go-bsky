package firehose

// Operation is the type for the 'Operation' field of [MessageHeader].
type Operation int64

const (
	OperationRegular Operation = 1
	OperationError   Operation = -1
)
