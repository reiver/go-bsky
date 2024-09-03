package graph

// Follow represents a 'app.bsky.graph.follow'.
type Follow struct {
	Type      string `json:"$type"`
	Subject   string `json:"subject"`
	CreatedAt string `json:"createdAt"`
}
