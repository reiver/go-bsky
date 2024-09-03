package feed

// Post represents a 'app.bsky.feed.post'.
type Post struct {
	Type      PostType `json:"$type"`
	CreatedAt string   `json:"createdAt"`
	Text      string   `json:"text"`
}

type PostType struct {}

func (PostType) MarshalText() ([]byte, error) {
	return []byte("app.bsky.feed.post"), nil
}
