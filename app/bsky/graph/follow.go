package graph

// Follow represents a 'app.bsky.graph.follow'.
type Follow struct {
	Type      string `json:"$type"`
	CreatedAt string `json:"createdAt"`
	Subject   string `json:"subject"`
}
