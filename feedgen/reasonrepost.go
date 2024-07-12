package feedgen

type internalReasonRepost struct {
	uri string
}

var _ Reason = internalReasonRepost{}

// ReasonReport returns a Reason that represents the optional reason in an item list from the "/xrpc/app.bsky.feed.getFeedSkeleton" response that is a repost.
//
// For example, if the reason is the JSON:
//
//	{
//		"$type":  "app.bsky.feed.defs#skeletonReasonRepost",
//		"repost": "at://did:example:1234/app.bsky.feed.post/9876543210"
//	}
func ReasonReport(uri string) Reason {
	return internalReasonRepost{
		uri:uri,
	}
}

func (receiver internalReasonRepost) ReasonAttributes() map[string]interface{} {
	return map[string]interface{}{
		"repost":receiver.uri,
	}
}

func (receiver internalReasonRepost) ReasonType() string {
	return "app.bsky.feed.defs#skeletonReasonRepost"
}
