package benchmarks

import "encoding/json"

type Event struct {
	Kind      int        `json:"kind"`
	CreatedAt int64      `json:"created_at"`
	Content   string     `json:"content"`
	PubKey    string     `json:"pubkey"`
	Sig       string     `json:"sig"`
	ID        string     `json:"id"`
	Tags      [][]string `json:"tags"`
}

func (e Event) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type EventShort struct {
	CreatedAt int    `json:"created_at"`
	Content   string `json:"content"`
	PubKey    string `json:"pubkey"`
}
