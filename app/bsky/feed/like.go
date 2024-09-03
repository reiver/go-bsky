package graph

// Like represents a 'app.bsky.feed.like'.
type Like struct {
	Type      string `json:"$type"`
	CreatedAt string `json:"createdAt"`
	Subject   struct{
		URI string `json:"uri"`
		CID string `json:"cid"`
	} `json:"subject"`
}
