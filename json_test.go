package benchmarks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/buger/jsonparser"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	gojson "github.com/goccy/go-json"
	"github.com/mailru/easyjson"
	"github.com/mreiferson/go-ujson"
	"github.com/nbd-wtf/go-nostr"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/tidwall/gjson"
)

// for each event we want to print the content, pubkey and timestamp

func BenchmarkShortEvent(b *testing.B) {
	sonic.Pretouch(reflect.TypeOf(EventShort{}))
	events := loadEvents()

	b.Run("json.Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event Event
				json.Unmarshal([]byte(evtstr), &event)
			}
		}
	})

	b.Run("gjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				v := gjson.Parse(evtstr)
				v.Get("content").String()
				v.Get("created_at").Int()
				v.Get("pubkey").String()
			}
		}
	})

	b.Run("jsonparser", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				b := []byte(evtstr)
				jsonparser.GetString(b, "content")
				jsonparser.GetInt(b, "created_at")
				jsonparser.GetUnsafeString(b, "pubkey")
			}
		}
	})

	b.Run("easyjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event EventShort
				easyjson.Unmarshal([]byte(evtstr), &event)
			}
		}
	})

	b.Run("ffjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			decoder := ffjson.NewDecoder()
			for _, evtstr := range events {
				var event EventShort
				decoder.DecodeFast([]byte(evtstr), &event)
			}
		}
	})

	b.Run("go-json", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event EventShort
				gojson.Unmarshal([]byte(evtstr), &event)
			}
		}
	})

	b.Run("ujson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				b := []byte(evtstr)
				uj, _ := ujson.NewFromBytes(b)
				uj.Get("content").MaybeString()
				uj.Get("created_at").MaybeInt64()
				uj.Get("pubkey").MaybeString()
			}
		}
	})

	b.Run("sonic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event EventShort
				sonic.UnmarshalString(evtstr, &event)
			}
		}
	})

	b.Run("sonic/get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				n1, _ := sonic.GetFromString(evtstr, "content")
				n1.String()
				n2, _ := sonic.GetFromString(evtstr, "created_at")
				n2.Int64()
				n3, _ := sonic.GetFromString(evtstr, "pubkey")
				n3.String()
			}
		}
	})

	b.Run("sonic/searcher/get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				s := ast.NewSearcher(evtstr)
				n1, _ := s.GetByPath("content")
				n1.String()
				n2, _ := s.GetByPath("created_at")
				n2.Int64()
				n3, _ := s.GetByPath("pubkey")
				n3.String()
			}
		}
	})
}

func BenchmarkLazyEvent(b *testing.B) {
	events := loadEvents()

	doStuff1 := func(l LazyEvent) {
		l.ID()
		l.Content()
		l.PubKey()
	}

	doStuff2 := func(l LazyEvent) {
		l.Content()
		l.Content()
		l.Content()
		l.Content()
		l.CreatedAt().Unix()
		l.CreatedAt().Unix()
		l.PubKey()
		l.PubKey()
		l.PubKey()
		l.PubKey()
	}

	doStuff3 := func(l LazyEvent) {
		l.Content()
		l.Content()
		l.ID()
		l.Content()
		l.Content()
		l.CreatedAt().Unix()
		l.CreatedAt().Unix()
		l.PubKey()
		l.PubKey()
		l.ID()
		l.PubKey()
		l.Content()
		l.Content()
		l.ID()
		l.Content()
		l.Content()
		l.CreatedAt().Unix()
		l.CreatedAt().Unix()
		l.PubKey()
		l.PubKey()
		l.ID()
		l.PubKey()
		l.ID()
		l.PubKey()
		l.PubKey()
	}

	actions := []func(LazyEvent){doStuff1, doStuff2, doStuff3}

	b.Run("gjson", func(b *testing.B) {
		for i, doStuff := range actions {
			b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, evtstr := range events {
						l := &LazyEventGJSON{v: gjson.Parse(evtstr)}
						doStuff(l)
					}
				}
			})
		}
	})

	b.Run("gjson-megalazy", func(b *testing.B) {
		for i, doStuff := range actions {
			b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, evtstr := range events {
						l := &LazyEventGJSON{v: gjson.Parse(evtstr)}
						doStuff(l)
					}
				}
			})
		}
	})

	b.Run("sonic", func(b *testing.B) {
		for i, doStuff := range actions {
			b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, evtstr := range events {
						l := &LazyEventSonic{searcher: ast.NewSearcher(evtstr)}
						doStuff(l)
					}
				}
			})
		}
	})

	b.Run("sonic-megalazy", func(b *testing.B) {
		for i, doStuff := range actions {
			b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, evtstr := range events {
						l := &MegaLazyEventSonic{searcher: ast.NewSearcher(evtstr)}
						doStuff(l)
					}
				}
			})
		}
	})

	b.Run("sonic-not-lazy", func(b *testing.B) {
		for i, doStuff := range actions {
			b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, evtstr := range events {
						var event Event
						sonic.UnmarshalString(evtstr, &event)
						l := &NotLazyEventSonic{embedded: event}
						doStuff(l)
					}
				}
			})
		}
	})
}

func BenchmarkFullEvent(b *testing.B) {
	sonic.Pretouch(reflect.TypeOf(Event{}))
	events := loadEvents()

	b.Run("json.Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event EventShort
				json.Unmarshal([]byte(evtstr), &event)
			}
		}
	})

	b.Run("gjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				v := gjson.Parse(evtstr)
				v.Get("kind").Int()
				v.Get("content").String()
				v.Get("created_at").Int()
				v.Get("pubkey").String()
				v.Get("tags").ForEach(func(_, v gjson.Result) bool {
					v.Get("1").String()
					return true
				})
				v.Get("sig").String()
				v.Get("id").String()
			}
		}
	})

	b.Run("gjson assign", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event Event
				v := gjson.Parse(evtstr)
				event.Kind = int(v.Get("kind").Int())
				event.Content = v.Get("content").String()
				event.CreatedAt = time.Unix(v.Get("created_at").Int(), 0)
				event.PubKey = v.Get("pubkey").String()
				v.Get("tags").ForEach(func(_, v gjson.Result) bool {
					tag := make([]string, 0, 4)
					v.ForEach(func(_, v gjson.Result) bool {
						tag = append(tag, v.String())
						return true
					})
					event.Tags = append(event.Tags, tag)
					return true
				})
				event.Sig = v.Get("sig").String()
				event.ID = v.Get("id").String()
			}
		}
	})

	b.Run("jsonparser", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				b := []byte(evtstr)
				jsonparser.GetInt(b, "kind")
				jsonparser.GetString(b, "content")
				jsonparser.GetInt(b, "created_at")
				jsonparser.GetUnsafeString(b, "pubkey")
				jsonparser.ArrayEach(b, func(tagb []byte, _ jsonparser.ValueType, _ int, _ error) {
					for i := 0; i < 10; i++ {
						_, err := jsonparser.GetUnsafeString(tagb, "[0]")
						if err != nil {
							break
						}
					}
				}, "tags")
				jsonparser.GetUnsafeString(b, "id")
				jsonparser.GetUnsafeString(b, "sig")
				jsonparser.GetInt(b, "kind")
			}
		}
	})

	b.Run("jsonparser assign", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event Event
				b := []byte(evtstr)
				kind, _ := jsonparser.GetInt(b, "kind")
				event.Kind = int(kind)
				event.Content, _ = jsonparser.GetString(b, "content")
				ts, _ := jsonparser.GetInt(b, "created_at")
				event.CreatedAt = time.Unix(ts, 0)
				event.PubKey, _ = jsonparser.GetUnsafeString(b, "pubkey")
				jsonparser.ArrayEach(b, func(tagb []byte, _ jsonparser.ValueType, _ int, _ error) {
					tag := make([]string, 0, 5)
					for i := 0; i < 10; i++ {
						v, err := jsonparser.GetUnsafeString(tagb, "[0]")
						if err != nil {
							break
						}
						tag = append(tag, v)
					}
					event.Tags = append(event.Tags, tag)
				}, "tags")
				event.ID, _ = jsonparser.GetUnsafeString(b, "id")
				event.Sig, _ = jsonparser.GetUnsafeString(b, "sig")
			}
		}
	})

	b.Run("easyjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event Event
				easyjson.Unmarshal([]byte(evtstr), &event)
			}
		}
	})

	b.Run("ffjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			decoder := ffjson.NewDecoder()
			for _, evtstr := range events {
				var event Event
				decoder.Decode([]byte(evtstr), &event)
			}
		}
	})

	b.Run("go-json", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event Event
				gojson.Unmarshal([]byte(evtstr), &event)
			}
		}
	})

	b.Run("ujson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				b := []byte(evtstr)
				uj, _ := ujson.NewFromBytes(b)
				uj.Get("kind").MaybeInt64()
				uj.Get("content").MaybeString()
				uj.Get("created_at").MaybeInt64()
				uj.Get("pubkey").MaybeString()
				uj.Get("sig").MaybeString()
				uj.Get("id").MaybeString()
				tagarr, _ := uj.Get("tags").MaybeArray()
				for _, tagi := range tagarr {
					itemarr, _ := tagi.MaybeArray()
					for _, item := range itemarr {
						item.MaybeString()
					}
				}
			}
		}
	})

	b.Run("sonic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event Event
				sonic.UnmarshalString(evtstr, &event)
			}
		}
	})
}

func BenchmarkGoNostrEventTyped(b *testing.B) {
	sonic.Pretouch(reflect.TypeOf(Event{}))
	api := sonic.Config{
		NoQuoteTextMarshaler: true,
		UseInt64:             true,
		NoNullSliceOrMap:     true,
	}.Froze()
	events := loadEvents()

	b.Run("go-nostr (fastjson)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event nostr.Event
				json.Unmarshal([]byte(evtstr), &event)
			}
		}
	})

	b.Run("sonic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event Event
				api.UnmarshalFromString(evtstr, &event)
			}
		}
	})

	b.Run("sonic/searcher/get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				s := ast.NewSearcher(evtstr)
				evt := Event{}
				id, _ := s.GetByPath("id")
				evt.ID, _ = id.StrictString()
				pubkey, _ := s.GetByPath("pubkey")
				evt.PubKey, _ = pubkey.StrictString()
				sig, _ := s.GetByPath("sig")
				evt.Sig, _ = sig.StrictString()
				createdAt, _ := s.GetByPath("created_at")
				createdAtInt64, _ := createdAt.StrictInt64()
				evt.CreatedAt = time.Unix(createdAtInt64, 0)
				content, _ := s.GetByPath("content")
				evt.Content, _ = content.StrictString()
				tagsv, _ := s.GetByPath("tags")
				evt.Tags = make(Tags, 0, 15)
				for i := 0; i < 40; i++ {
					tagv := tagsv.Index(i)
					if tagv.Exists() {
						tag := make(Tag, 0, 10)
						for j := 0; j < 10; j++ {
							itemv := tagv.Index(j)
							v, _ := itemv.StrictString()
							tag = append(tag, v)
						}
						evt.Tags = append(evt.Tags, tag)
					}
				}
			}
		}
	})

	b.Run("easyjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event Event
				easyjson.Unmarshal([]byte(evtstr), &event)
			}
		}
	})
}

func BenchmarkEnvelope(b *testing.B) {
	sonic.Pretouch(reflect.TypeOf(Event{}))
	api := sonic.Config{
		NoQuoteTextMarshaler: true,
		UseInt64:             true,
		NoNullSliceOrMap:     true,
	}.Froze()
	lines := loadLines()

	b.Run("json.Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, line := range lines {
				env := make([]json.RawMessage, 0, 4)
				json.Unmarshal([]byte(line), &env)
				if len(env) == 0 {
					continue
				}
				var typ string
				json.Unmarshal(env[0], &typ)
				switch typ {
				case "EVENT":
					if len(env) != 2 {
						continue
					}
					var sub string
					json.Unmarshal(env[1], &sub)
					var evt Event
					json.Unmarshal(env[2], &evt)
				case "OK":
					if len(env) < 3 {
						continue
					}
					var id string
					var ok bool
					json.Unmarshal(env[1], &id)
					json.Unmarshal(env[2], &ok)
					if len(env) > 3 {
						var msg string
						json.Unmarshal(env[3], &msg)
					}
				case "EOSE":
					if len(env) != 2 {
						continue
					}
					var sub string
					json.Unmarshal(env[1], &sub)
				case "NOTICE":
					if len(env) != 2 {
						continue
					}
					var msg string
					json.Unmarshal(env[1], &msg)
				}
			}
		}
	})

	b.Run("go-nostr (fastjson)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, line := range lines {
				env := make([]json.RawMessage, 0, 4)
				json.Unmarshal([]byte(line), &env)
				if len(env) == 0 {
					continue
				}
				var typ string
				json.Unmarshal(env[0], &typ)
				switch typ {
				case "EVENT":
					if len(env) != 2 {
						continue
					}
					var sub string
					json.Unmarshal(env[1], &sub)
					var evt nostr.Event
					json.Unmarshal(env[2], &evt)
				case "OK":
					if len(env) < 3 {
						continue
					}
					var id string
					var ok bool
					json.Unmarshal(env[1], &id)
					json.Unmarshal(env[2], &ok)
					if len(env) > 3 {
						var msg string
						json.Unmarshal(env[3], &msg)
					}
				case "EOSE":
					if len(env) != 2 {
						continue
					}
					var sub string
					json.Unmarshal(env[1], &sub)
				case "NOTICE":
					if len(env) != 2 {
						continue
					}
					var msg string
					json.Unmarshal(env[1], &msg)
				}
			}
		}
	})

	b.Run("sonic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, line := range lines {
				s := ast.NewSearcher(line)
				first, _ := s.GetByPath(0)

				cmd, _ := first.StrictString()

				switch cmd {
				case "EVENT":
					subv, _ := s.GetByPath(1)
					subv.StrictString()
					evtv, _ := s.GetByPath(2)
					var evt Event
					raw, _ := evtv.Raw()
					api.UnmarshalFromString(raw, &evt)
				case "OK":
					idv, _ := s.GetByPath(1)
					okv, _ := s.GetByPath(2)
					msgv, _ := s.GetByPath(3)
					idv.StrictString()
					okv.Bool()
					msgv.StrictString()
				case "EOSE":
					subv, _ := s.GetByPath(1)
					subv.StrictString()
				case "NOTICE":
					msgv, _ := s.GetByPath(1)
					msgv.StrictString()
				}
			}
		}
	})

	b.Run("easyjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, line := range lines {
				env := make([]json.RawMessage, 0, 4)
				json.Unmarshal([]byte(line), &env)
				if len(env) == 0 {
					continue
				}
				var typ string
				json.Unmarshal(env[0], &typ)
				switch typ {
				case "EVENT":
					if len(env) != 2 {
						continue
					}
					var sub string
					json.Unmarshal(env[1], &sub)
					var event Event
					easyjson.Unmarshal(env[2], &event)
				case "OK":
					if len(env) < 3 {
						continue
					}
					var id string
					var ok bool
					json.Unmarshal(env[1], &id)
					json.Unmarshal(env[2], &ok)
					if len(env) > 3 {
						var msg string
						json.Unmarshal(env[3], &msg)
					}
				case "EOSE":
					if len(env) != 2 {
						continue
					}
					var sub string
					json.Unmarshal(env[1], &sub)
				case "NOTICE":
					if len(env) != 2 {
						continue
					}
					var msg string
					json.Unmarshal(env[1], &msg)
				}
			}
		}
	})

	b.Run("gjson + easyjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, line := range lines {
				v := gjson.Parse(line)
				switch v.Get("0").Str {
				case "EVENT":
					_ = v.Get("1").Str
					var event Event
					easyjson.Unmarshal([]byte(v.Get("2").Raw), &event)
				case "OK":
					_ = v.Get("1").Str
					_ = v.Get("2").Bool()
					_ = v.Get("3").Str
				case "EOSE":
					_ = v.Get("1").Str
				case "NOTICE":
					_ = v.Get("1").Str
				}
			}
		}
	})
}

func loadLines() []string {
	b, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	text := string(b)
	lines := strings.Split(text, "\n")
	return lines[0 : len(lines)-1]
}

func loadEvents() []string {
	lines := loadLines()
	events := make([]string, 0, len(lines))

	for _, line := range lines {
		if strings.HasPrefix(line, "[\"EVENT") {
			events = append(events, line[13:len(line)-1])
		}
	}

	return events
}
