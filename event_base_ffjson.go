// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: event_base.go

package benchmarks

import (
	"bytes"
	"encoding/json"
	"fmt"

	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSONBuf marshal buff to json - template
func (j *Event) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"kind":`)
	fflib.FormatBits2(buf, uint64(j.Kind), 10, j.Kind < 0)
	buf.WriteString(`,"created_at":`)
	fflib.FormatBits2(buf, uint64(j.CreatedAt), 10, j.CreatedAt < 0)
	buf.WriteString(`,"content":`)
	fflib.WriteJsonString(buf, string(j.Content))
	buf.WriteString(`,"pubkey":`)
	fflib.WriteJsonString(buf, string(j.PubKey))
	buf.WriteString(`,"sig":`)
	fflib.WriteJsonString(buf, string(j.Sig))
	buf.WriteString(`,"id":`)
	fflib.WriteJsonString(buf, string(j.ID))
	buf.WriteString(`,"tags":`)
	if j.Tags != nil {
		buf.WriteString(`[`)
		for i, v := range j.Tags {
			if i != 0 {
				buf.WriteString(`,`)
			}
			if v != nil {
				buf.WriteString(`[`)
				for i, v := range v {
					if i != 0 {
						buf.WriteString(`,`)
					}
					fflib.WriteJsonString(buf, string(v))
				}
				buf.WriteString(`]`)
			} else {
				buf.WriteString(`null`)
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtEventbase = iota
	ffjtEventnosuchkey

	ffjtEventKind

	ffjtEventCreatedAt

	ffjtEventContent

	ffjtEventPubKey

	ffjtEventSig

	ffjtEventID

	ffjtEventTags
)

var ffjKeyEventKind = []byte("kind")

var ffjKeyEventCreatedAt = []byte("created_at")

var ffjKeyEventContent = []byte("content")

var ffjKeyEventPubKey = []byte("pubkey")

var ffjKeyEventSig = []byte("sig")

var ffjKeyEventID = []byte("id")

var ffjKeyEventTags = []byte("tags")

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Event) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtEventbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtEventnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffjKeyEventCreatedAt, kn) {
						currentKey = ffjtEventCreatedAt
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyEventContent, kn) {
						currentKey = ffjtEventContent
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyEventID, kn) {
						currentKey = ffjtEventID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'k':

					if bytes.Equal(ffjKeyEventKind, kn) {
						currentKey = ffjtEventKind
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyEventPubKey, kn) {
						currentKey = ffjtEventPubKey
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyEventSig, kn) {
						currentKey = ffjtEventSig
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyEventTags, kn) {
						currentKey = ffjtEventTags
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyEventTags, kn) {
					currentKey = ffjtEventTags
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyEventID, kn) {
					currentKey = ffjtEventID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyEventSig, kn) {
					currentKey = ffjtEventSig
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyEventPubKey, kn) {
					currentKey = ffjtEventPubKey
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyEventContent, kn) {
					currentKey = ffjtEventContent
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyEventCreatedAt, kn) {
					currentKey = ffjtEventCreatedAt
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyEventKind, kn) {
					currentKey = ffjtEventKind
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtEventnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtEventKind:
					goto handle_Kind

				case ffjtEventCreatedAt:
					goto handle_CreatedAt

				case ffjtEventContent:
					goto handle_Content

				case ffjtEventPubKey:
					goto handle_PubKey

				case ffjtEventSig:
					goto handle_Sig

				case ffjtEventID:
					goto handle_ID

				case ffjtEventTags:
					goto handle_Tags

				case ffjtEventnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Kind:

	/* handler: j.Kind type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {
		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)
			if err != nil {
				return fs.WrapErr(err)
			}

			j.Kind = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CreatedAt:

	/* handler: j.CreatedAt type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {
		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)
			if err != nil {
				return fs.WrapErr(err)
			}

			j.CreatedAt = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Content:

	/* handler: j.Content type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {
		} else {

			outBuf := fs.Output.Bytes()

			j.Content = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PubKey:

	/* handler: j.PubKey type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {
		} else {

			outBuf := fs.Output.Bytes()

			j.PubKey = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Sig:

	/* handler: j.Sig type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {
		} else {

			outBuf := fs.Output.Bytes()

			j.Sig = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ID:

	/* handler: j.ID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {
		} else {

			outBuf := fs.Output.Bytes()

			j.ID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Tags:

	/* handler: j.Tags type=[][]string kind=slice quoted=false*/

	{
		/* Falling back. type=[][]string kind=slice */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Tags)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *EventShort) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *EventShort) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"created_at":`)
	fflib.FormatBits2(buf, uint64(j.CreatedAt), 10, j.CreatedAt < 0)
	buf.WriteString(`,"content":`)
	fflib.WriteJsonString(buf, string(j.Content))
	buf.WriteString(`,"pubkey":`)
	fflib.WriteJsonString(buf, string(j.PubKey))
	buf.WriteByte('}')
	return nil
}

const (
	ffjtEventShortbase = iota
	ffjtEventShortnosuchkey

	ffjtEventShortCreatedAt

	ffjtEventShortContent

	ffjtEventShortPubKey
)

var ffjKeyEventShortCreatedAt = []byte("created_at")

var ffjKeyEventShortContent = []byte("content")

var ffjKeyEventShortPubKey = []byte("pubkey")

// UnmarshalJSON umarshall json - template of ffjson
func (j *EventShort) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *EventShort) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtEventShortbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtEventShortnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffjKeyEventShortCreatedAt, kn) {
						currentKey = ffjtEventShortCreatedAt
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyEventShortContent, kn) {
						currentKey = ffjtEventShortContent
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyEventShortPubKey, kn) {
						currentKey = ffjtEventShortPubKey
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyEventShortPubKey, kn) {
					currentKey = ffjtEventShortPubKey
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyEventShortContent, kn) {
					currentKey = ffjtEventShortContent
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyEventShortCreatedAt, kn) {
					currentKey = ffjtEventShortCreatedAt
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtEventShortnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtEventShortCreatedAt:
					goto handle_CreatedAt

				case ffjtEventShortContent:
					goto handle_Content

				case ffjtEventShortPubKey:
					goto handle_PubKey

				case ffjtEventShortnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_CreatedAt:

	/* handler: j.CreatedAt type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {
		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)
			if err != nil {
				return fs.WrapErr(err)
			}

			j.CreatedAt = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Content:

	/* handler: j.Content type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {
		} else {

			outBuf := fs.Output.Bytes()

			j.Content = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PubKey:

	/* handler: j.PubKey type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {
		} else {

			outBuf := fs.Output.Bytes()

			j.PubKey = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
