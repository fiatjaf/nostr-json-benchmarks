package benchmarks

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

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
				sonic.Unmarshal([]byte(evtstr), &event)
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

func BenchmarkFullEvent(b *testing.B) {
	events := loadEvents()

	b.Run("json.Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event EventShort
				json.Unmarshal([]byte(evtstr), &event)
			}
		}
	})

	b.Run("go-nostr (fastjson)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtstr := range events {
				var event nostr.Event
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
				event.CreatedAt = v.Get("created_at").Int()
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
				event.CreatedAt, _ = jsonparser.GetInt(b, "created_at")
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
				sonic.Unmarshal([]byte(evtstr), &event)
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

	for i, line := range lines {
		lines[i] = line[13 : len(line)-1]
	}

	return lines
}
