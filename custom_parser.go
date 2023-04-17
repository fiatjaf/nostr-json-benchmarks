package benchmarks

import (
	"fmt"
	"strconv"
	"strings"
)

func parseEvent(j string) (*Event, error) {
	evt := Event{
		Tags: make(Tags, 0, 20),

		extra: make(map[string]any),
	}

	missingFields := 7

	index := strings.IndexByte(j, byte('{'))
	j = j[index:]

	var key string
	for {
		// key
		startKey := strings.IndexByte(j, byte('"'))
		j = j[startKey+1:]
		endKey := strings.IndexByte(j, byte('"'))
		key = j[0:endKey]
		j = j[endKey+1:]

		// colon
		colon := strings.IndexByte(j, byte(':'))
		j = j[colon+1:]

		// value
		switch key {
		case "content":
			missingFields--
			// expect a string with escape sequences in it
			startStr := strings.IndexByte(j, byte('"'))
			j = j[startStr+1:]

			endStr := strings.IndexByte(j, byte('"'))
			quoted := false
			if endStr > 0 {
				for j[endStr-1] == '\\' {
					endStr = endStr + 1 + strings.IndexByte(j[endStr+1:], byte('"'))
					quoted = true
				}
			}
			str := j[0:endStr]
			if quoted {
				evt.Content, _ = strconv.Unquote(str)
			} else {
				evt.Content = str
			}
			j = j[endStr+1:]
		case "tags":
			missingFields--
			// expect an array of arrays of strings
			startTopArray := strings.IndexByte(j, byte('['))
			j = j[startTopArray+1:]
			for {
				startTag := strings.IndexByte(j, byte('['))
				if startTag == -1 {
					// if there are no more "["s we're at the end of the json event, so we must ensure endTopArray is found
					startTag = len(j)
				}
				endTopArray := strings.IndexByte(j[0:startTag], byte(']')) // only search up to startTag
				if endTopArray != -1 && endTopArray < startTag {
					// no more tags
					j = j[endTopArray+1:]
					break
				}

				// there is a tag
				tag := make(Tag, 0, 5)
				j = j[startTag+1:]

				for {
					// expect a string
					startStr := strings.IndexByte(j, byte('"'))
					j = j[startStr+1:]

					endStr := strings.IndexByte(j, byte('"'))
					quoted := false
					if endStr > 0 {
						for j[endStr-1] == '\\' {
							endStr = endStr + 1 + strings.IndexByte(j[endStr+1:], byte('"'))
							quoted = true
						}
					}
					str := j[0:endStr]
					if quoted {
						evt.Content, _ = strconv.Unquote(str)
					} else {
						evt.Content = str
					}
					j = j[endStr+1:]

					// a comma or the tag end
					comma := strings.IndexByte(j, byte(','))
					if comma == -1 {
						// if there are no more commas we're at the end of the json event, so we must ensure endTag is found
						comma = len(j)
					}
					endTag := strings.IndexByte(j[0:comma], byte(']')) // only search up to comma
					if endTag != -1 && endTag < comma {
						evt.Tags = append(evt.Tags, tag)
						j = j[endTag+1:]
						break
					}
					j = j[comma+1:]
				}
			}

		case "id":
			missingFields--
			// expect 64 characters hex
			startStr := strings.IndexByte(j, byte('"'))
			j = j[startStr+1:]
			evt.ID = j[0:64]
			j = j[64+1:]
		case "pubkey":
			missingFields--
			// expect 64 characters hex
			startStr := strings.IndexByte(j, byte('"'))
			j = j[startStr+1:]
			evt.PubKey = j[0:64]
			j = j[64+1:]
		case "sig":
			missingFields--
			// expect 128 characters hex
			startStr := strings.IndexByte(j, byte('"'))
			j = j[startStr+1:]
			evt.Sig = j[0:128]
			j = j[128+1:]
		case "created_at":
			missingFields--
			// expect a bunch of number characters
			endNumber := strings.IndexByte(j, byte(','))
			if endNumber == -1 && missingFields == 1 {
				endNumber = strings.IndexByte(j, byte('}'))
				if endNumber == -1 {
					return nil, fmt.Errorf("failed to parse created_at: %s", j[0:11])
				}
			}
			number, err := strconv.ParseUint(strings.TrimSpace(j[0:endNumber]), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("created_at is not a number: %w", err)
			}
			evt.CreatedAt = Timestamp(number)
		case "kind":
			missingFields--
			// expect a bunch of number characters
			endNumber := strings.IndexByte(j, byte(','))
			if endNumber == -1 && missingFields == 1 {
				endNumber = strings.IndexByte(j, byte('}'))
				if endNumber == -1 {
					return nil, fmt.Errorf("failed to parse kind: '%s'", j[0:6])
				}
			}
			number, err := strconv.ParseUint(strings.TrimSpace(j[0:endNumber]), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("kind is not a number: '%s': %w", j[0:6], err)
			}
			evt.Kind = int(number)
		}

		if missingFields == 0 {
			break
		}
	}

	return &evt, nil
}

var sequence = []byte{'{', '"', ':'}
