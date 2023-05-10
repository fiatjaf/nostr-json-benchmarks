package benchmarks

import "time"

type LazyEventNson struct {
	evtstr string
}

func (l *LazyEventNson) Kind() int            { return nsonGetKind(l.evtstr) }
func (l *LazyEventNson) CreatedAt() time.Time { return nsonGetCreatedAt(l.evtstr).Time() }
func (l *LazyEventNson) Content() string      { return nsonGetContent(l.evtstr) }
func (l *LazyEventNson) ID() string           { return nsonGetID(l.evtstr) }
func (l *LazyEventNson) PubKey() string       { return nsonGetPubkey(l.evtstr) }
func (l *LazyEventNson) Sig() string          { return nsonGetSig(l.evtstr) }
