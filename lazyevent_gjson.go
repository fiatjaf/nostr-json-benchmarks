package benchmarks

import (
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

type LazyEventGJSON struct {
	v gjson.Result

	kind      *int
	createdAt *time.Time
	content   *string
	pubkey    *string
	sig       *string
	id        *string
	tags      *[][]string
}

func (l *LazyEventGJSON) Kind() int {
	if l.kind != nil {
		return *l.kind
	}
	v := l.v.Get("kind").Int()
	kind := int(v)
	l.kind = &kind
	return kind
}

func (l *LazyEventGJSON) CreatedAt() time.Time {
	if l.createdAt != nil {
		return *l.createdAt
	}
	raw := l.v.Get("created_at").Raw

	// since this is never going to be negative or a float, parse naïvely
	// TODO: improve parsing here by ignoring all edge cases covered by strconv
	i64, _ := strconv.ParseUint(raw, 10, 64)

	createdAt := time.Unix(int64(i64), 0)
	l.createdAt = &createdAt
	return createdAt
}

func (l *LazyEventGJSON) Content() string {
	if l.content != nil {
		return *l.content
	}
	v := l.v.Get("content").Str
	l.content = &v
	return v
}

func (l *LazyEventGJSON) ID() string {
	if l.id != nil {
		return *l.id
	}
	v := l.v.Get("id").Str
	l.id = &v
	return v
}

func (l *LazyEventGJSON) PubKey() string {
	if l.pubkey != nil {
		return *l.pubkey
	}
	v := l.v.Get("pubkey").Str
	l.pubkey = &v
	return v
}

func (l *LazyEventGJSON) Sig() string {
	if l.sig != nil {
		return *l.sig
	}
	v := l.v.Get("sig").Str
	l.sig = &v
	return v
}

type MegaLazyEventGJSON struct {
	v gjson.Result
}

func (l *MegaLazyEventGJSON) Kind() int {
	v := l.v.Get("kind").Int()
	kind := int(v)
	return kind
}

func (l *MegaLazyEventGJSON) CreatedAt() time.Time {
	raw := l.v.Get("created_at").Raw

	// since this is never going to be negative or a float, parse naïvely
	// TODO: improve parsing here by ignoring all edge cases covered by strconv
	i64, _ := strconv.ParseUint(raw, 10, 64)

	createdAt := time.Unix(int64(i64), 0)
	return createdAt
}

func (l *MegaLazyEventGJSON) Content() string {
	v := l.v.Get("content").Str
	return v
}

func (l *MegaLazyEventGJSON) ID() string {
	v := l.v.Get("id").Str
	return v
}

func (l *MegaLazyEventGJSON) PubKey() string {
	v := l.v.Get("pubkey").Str
	return v
}

func (l *MegaLazyEventGJSON) Sig() string {
	v := l.v.Get("sig").Str
	return v
}
