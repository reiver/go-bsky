package feed

// Repost represents a 'app.bsky.feed.repost'.
type Repost struct {
	Type      RepostType `json:"$type"`
	CreatedAt string     `json:"createdAt"`
	Subject   struct{
		URI string   `json:"uri"`
		CID string   `json:"cid"`
	}                    `json:"subject"`
}

type RepostType struct {}

func (RepostType) MarshalText() ([]byte, error) {
	return []byte("app.bsky.feed.repost"), nil
}
