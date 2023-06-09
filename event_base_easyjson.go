// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package benchmarks

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson723fdcd4DecodeGithubComFiatjafNostrJsonBenchmarks(in *jlexer.Lexer, out *EventShort) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "created_at":
			out.CreatedAt = int(in.Int())
		case "content":
			out.Content = string(in.String())
		case "pubkey":
			out.PubKey = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson723fdcd4EncodeGithubComFiatjafNostrJsonBenchmarks(out *jwriter.Writer, in EventShort) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix[1:])
		out.Int(int(in.CreatedAt))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	{
		const prefix string = ",\"pubkey\":"
		out.RawString(prefix)
		out.String(string(in.PubKey))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventShort) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson723fdcd4EncodeGithubComFiatjafNostrJsonBenchmarks(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventShort) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson723fdcd4DecodeGithubComFiatjafNostrJsonBenchmarks(l, v)
}
func easyjson723fdcd4DecodeGithubComFiatjafNostrJsonBenchmarks1(in *jlexer.Lexer, out *Event) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "kind":
			out.Kind = int(in.Int())
		case "created_at":
			out.CreatedAt = Timestamp(in.Int64())
		case "content":
			out.Content = string(in.String())
		case "pubkey":
			out.PubKey = string(in.String())
		case "sig":
			out.Sig = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make(Tags, 0, 2)
					} else {
						out.Tags = Tags{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Tag
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						in.Delim('[')
						if v1 == nil {
							if !in.IsDelim(']') {
								v1 = make(Tag, 0, 4)
							} else {
								v1 = Tag{}
							}
						} else {
							v1 = (v1)[:0]
						}
						for !in.IsDelim(']') {
							var v2 string
							v2 = string(in.String())
							v1 = append(v1, v2)
							in.WantComma()
						}
						in.Delim(']')
					}
					out.Tags = append(out.Tags, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson723fdcd4EncodeGithubComFiatjafNostrJsonBenchmarks1(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Kind))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Int64(int64(in.CreatedAt))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	{
		const prefix string = ",\"pubkey\":"
		out.RawString(prefix)
		out.String(string(in.PubKey))
	}
	{
		const prefix string = ",\"sig\":"
		out.RawString(prefix)
		out.String(string(in.Sig))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Tags {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v5, v6 := range v4 {
						if v5 > 0 {
							out.RawByte(',')
						}
						out.String(string(v6))
					}
					out.RawByte(']')
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson723fdcd4EncodeGithubComFiatjafNostrJsonBenchmarks1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson723fdcd4DecodeGithubComFiatjafNostrJsonBenchmarks1(l, v)
}
