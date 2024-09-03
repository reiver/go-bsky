package graph

// Post represents a 'app.bsky.feed.post'.
type Post struct {
	Type      string `json:"$type"`
	CreatedAt string `json:"createdAt"`
	Text      string `json:"text"`
}
