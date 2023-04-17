package benchmarks

import (
	"encoding/json"
	"time"
)

type Event struct {
	Kind      int       `json:"kind"`
	CreatedAt Timestamp `json:"created_at"`
	Content   string    `json:"content"`
	PubKey    string    `json:"pubkey"`
	Sig       string    `json:"sig"`
	ID        string    `json:"id"`
	Tags      Tags      `json:"tags"`

	extra map[string]any
}

func (e Event) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type Timestamp int64

func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}

type EventShort struct {
	CreatedAt int    `json:"created_at"`
	Content   string `json:"content"`
	PubKey    string `json:"pubkey"`
}

type LazyEvent interface {
	Kind() int
	CreatedAt() time.Time
	Content() string
	PubKey() string
	Sig() string
	ID() string
	// Tags() [][]string
}

type (
	Tags []Tag
	Tag  []string
)
