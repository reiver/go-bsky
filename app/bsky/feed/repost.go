package graph

// Repost represents a 'app.bsky.feed.repost'.
type Repost struct {
	Type      string `json:"$type"`
	CreatedAt string `json:"createdAt"`
	Subject   struct{
		URI string `json:"uri"`
		CID string `json:"cid"`
	} `json:"subject"`
}
