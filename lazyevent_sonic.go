package benchmarks

import (
	"time"

	"github.com/bytedance/sonic/ast"
)

type LazyEventSonic struct {
	searcher *ast.Searcher

	kind      *int
	createdAt *time.Time
	content   *string
	pubkey    *string
	sig       *string
	id        *string
	tags      *[][]string
}

func (l *LazyEventSonic) Kind() int {
	if l.kind != nil {
		return *l.kind
	}
	v, _ := l.searcher.GetByPath("kind")
	i64, _ := v.Int64()
	kind := int(i64)
	l.kind = &kind
	return kind
}

func (l *LazyEventSonic) CreatedAt() time.Time {
	if l.createdAt != nil {
		return *l.createdAt
	}
	v, _ := l.searcher.GetByPath("created_at")
	i64, _ := v.Int64()
	createdAt := time.Unix(i64, 0)
	l.createdAt = &createdAt
	return createdAt
}

func (l *LazyEventSonic) Content() string {
	if l.content != nil {
		return *l.content
	}
	v, _ := l.searcher.GetByPath("content")
	str, _ := v.String()
	l.content = &str
	return str
}

func (l *LazyEventSonic) ID() string {
	if l.id != nil {
		return *l.id
	}
	v, _ := l.searcher.GetByPath("id")
	str, _ := v.String()
	l.id = &str
	return str
}

func (l *LazyEventSonic) PubKey() string {
	if l.pubkey != nil {
		return *l.pubkey
	}
	v, _ := l.searcher.GetByPath("pubkey")
	str, _ := v.String()
	l.pubkey = &str
	return str
}

func (l *LazyEventSonic) Sig() string {
	if l.sig != nil {
		return *l.sig
	}
	v, _ := l.searcher.GetByPath("sig")
	str, _ := v.String()
	l.sig = &str
	return str
}

type MegaLazyEventSonic struct {
	searcher *ast.Searcher
}

func (l *MegaLazyEventSonic) Kind() int {
	v, _ := l.searcher.GetByPath("kind")
	i64, _ := v.Int64()
	kind := int(i64)
	return kind
}

func (l *MegaLazyEventSonic) CreatedAt() time.Time {
	v, _ := l.searcher.GetByPath("created_at")
	i64, _ := v.Int64()
	createdAt := time.Unix(i64, 0)
	return createdAt
}

func (l *MegaLazyEventSonic) Content() string {
	v, _ := l.searcher.GetByPath("content")
	str, _ := v.String()
	return str
}

func (l *MegaLazyEventSonic) ID() string {
	v, _ := l.searcher.GetByPath("id")
	str, _ := v.String()
	return str
}

func (l *MegaLazyEventSonic) PubKey() string {
	v, _ := l.searcher.GetByPath("pubkey")
	str, _ := v.String()
	return str
}

func (l *MegaLazyEventSonic) Sig() string {
	v, _ := l.searcher.GetByPath("sig")
	str, _ := v.String()
	return str
}

type NotLazyEventSonic struct {
	embedded GoNostrEvent
}

func (l *NotLazyEventSonic) Kind() int {
	return l.embedded.Kind
}

func (l *NotLazyEventSonic) CreatedAt() time.Time {
	return l.embedded.CreatedAt
}

func (l *NotLazyEventSonic) Content() string {
	return l.embedded.Content
}

func (l *NotLazyEventSonic) ID() string {
	return l.embedded.ID
}

func (l *NotLazyEventSonic) PubKey() string {
	return l.embedded.PubKey
}

func (l *NotLazyEventSonic) Sig() string {
	return l.embedded.Sig
}
