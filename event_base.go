package benchmarks

import (
	"encoding/hex"
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

func (e Event) ToBinary() *EventBinary {
	eb := &EventBinary{
		Kind:      e.Kind,
		CreatedAt: e.CreatedAt,
		Content:   e.Content,
		Tags:      e.Tags,
	}
	eb.ID, _ = hex.DecodeString(e.ID)
	eb.PubKey, _ = hex.DecodeString(e.PubKey)
	eb.Sig, _ = hex.DecodeString(e.Sig)
	return eb
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

type EventBinary struct {
	Kind      int
	CreatedAt Timestamp
	Content   string
	PubKey    []byte
	Sig       []byte
	ID        []byte
	Tags      Tags
}

func (e EventBinary) ToNormal() *Event {
	en := &Event{
		Kind:      e.Kind,
		CreatedAt: e.CreatedAt,
		Content:   e.Content,
		Tags:      e.Tags,
	}
	en.ID = hex.EncodeToString(e.ID)
	en.PubKey = hex.EncodeToString(e.PubKey)
	en.Sig = hex.EncodeToString(e.Sig)
	return en
}
