package feedgen

// Reason represents the optional reason in an item list from the "/xrpc/app.bsky.feed.getFeedSkeleton" response.
//
// For example, if the reason is the JSON:
//
//	{
//		"$type":  "app.bsky.feed.defs#skeletonReasonRepost",
//		"repost": "at://did:example:1234/app.bsky.feed.post/9876543210"
//	}
//
// Then ReasonType() would return:
//
//		"app.bsky.feed.defs#skeletonReasonRepost"
//
// And ReasonAttributes() would return:
//
//	map[string]interface{}{
//		"repost": "at://did:example:1234/app.bsky.feed.post/9876543210",
//	}
type Reason interface {
	ReasonType() string
	ReasonAttributes() map[string]interface{}
}
