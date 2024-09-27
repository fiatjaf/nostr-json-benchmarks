package benchmarks

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
	"github.com/mailru/easyjson"
	"github.com/nbd-wtf/go-nostr"
	"github.com/tidwall/gjson"
)

func BenchmarkFullEvent(b *testing.B) {
	sonic.Pretouch(reflect.TypeOf(nostr.Event{}))
	events := loadEvents()

	b.Run("json.Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtj := range events {
				var event nostr.Event
				_ = json.Unmarshal([]byte(evtj), &event)
			}
		}
	})

	b.Run("easyjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtj := range events {
				var event nostr.Event
				_ = easyjson.Unmarshal([]byte(evtj), &event)
			}
		}
	})

	b.Run("go-json", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtj := range events {
				var event nostr.Event
				_ = gojson.Unmarshal([]byte(evtj), &event)
			}
		}
	})

	b.Run("sonic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, evtj := range events {
				var event nostr.Event
				_ = sonic.Unmarshal(evtj, &event)
			}
		}
	})
}

func BenchmarkEnvelope(b *testing.B) {
	sonic.Pretouch(reflect.TypeOf(nostr.Event{}))
	api := sonic.Config{
		NoQuoteTextMarshaler: true,
		UseInt64:             true,
		NoNullSliceOrMap:     true,
	}.Froze()
	lines := loadEnvelopes()

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
					if len(env) != 3 {
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

	b.Run("go-nostr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, line := range lines {
				_ = nostr.ParseMessage(line)
			}
		}
	})

	b.Run("sonic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, line := range lines {
				s, _ := sonic.Get(line)
				first := s.Index(0)

				cmd, _ := first.StrictString()

				switch cmd {
				case "EVENT":
					_, _ = s.Index(1).StrictString()
					evtv := s.Index(2)
					var evt nostr.Event
					raw, _ := evtv.Raw()
					api.UnmarshalFromString(raw, &evt)
				case "OK":
					_, _ = s.Index(1).StrictString()
					_, _ = s.Index(2).StrictString()

					v := s.Index(3)
					if v != nil {
						_, _ = v.StrictString()
					}
				case "EOSE":
					_, _ = s.Index(1).StrictString()
				case "NOTICE":
					_, _ = s.Index(1).StrictString()
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
					if len(env) != 3 {
						continue
					}
					var sub string
					json.Unmarshal(env[1], &sub)
					var event nostr.Event
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
				v := gjson.ParseBytes(line)
				switch v.Get("0").Str {
				case "EVENT":
					_ = v.Get("1").Str
					var event nostr.Event
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

	b.Run("gjson + sonic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, line := range lines {
				v := gjson.Parse(string(line))
				switch v.Get("0").Str {
				case "EVENT":
					_ = v.Get("1").Str
					var event nostr.Event
					api.UnmarshalFromString(v.Get("2").Raw, &event)
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

	b.Run("sonic + easyjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, line := range lines {
				s, _ := sonic.Get(line)
				first, _ := s.Index(0).StrictString()
				switch first {
				case "EVENT":
					_, _ = s.Index(1).StrictString()
					third, _ := s.Index(2).Raw()
					var event nostr.Event
					easyjson.Unmarshal([]byte(third), &event)
				case "OK":
					_, _ = s.Index(1).StrictString()
					_, _ = s.Index(2).Bool()

					v := s.Index(3)
					if v != nil {
						_, _ = v.StrictString()
					}
				case "EOSE":
					_, _ = s.Index(1).StrictString()
				case "NOTICE":
					_, _ = s.Index(1).StrictString()
				}
			}
		}
	})
}
