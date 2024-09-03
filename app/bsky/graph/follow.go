package graph

// Follow represents 'app.bsky.graph.follow' in the 
type Follow struct {
	Type      string `json:"$type"`
	Subject   string `json:"subject"`
	CreatedAt string `json:"createdAt"`
}
