package graph

// Follow represents a 'app.bsky.graph.follow'.
type Follow struct {
	Type      FollowType `json:"$type"`
	CreatedAt string     `json:"createdAt"`
	Subject   string     `json:"subject"`
}

type FollowType struct {}

func (FollowType) MarshalText() ([]byte, error) {
	return []byte("app.bsky.graph.follow"), nil
}
