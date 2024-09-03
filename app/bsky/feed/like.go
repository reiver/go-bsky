package feed

// Like represents a 'app.bsky.feed.like'.
type Like struct {
	Type      LikeType `json:"$type"`
	CreatedAt string   `json:"createdAt"`
	Subject   struct{
		URI string `json:"uri"`
		CID string `json:"cid"`
	}                  `json:"subject"`
}

type LikeType struct {}

func (LikeType) MarshalText() ([]byte, error) {
	return []byte("app.bsky.feed.like"), nil
}
