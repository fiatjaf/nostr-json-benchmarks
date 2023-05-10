package benchmarks

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/nbd-wtf/go-nostr"
)

func checkParsedCorrectly(t *testing.T, jevt string, evt *Event) {
	var canonical nostr.Event
	json.Unmarshal([]byte(jevt), &canonical)

	if evt.ID != canonical.ID {
		t.Errorf("id is wrong: %s != %s", evt.ID, canonical.ID)
	}
	if evt.PubKey != canonical.PubKey {
		t.Errorf("pubkey is wrong: %s != %s", evt.PubKey, canonical.PubKey)
	}
	if evt.Sig != canonical.Sig {
		t.Errorf("sig is wrong: %s != %s", evt.Sig, canonical.Sig)
	}
	if evt.Content != canonical.Content {
		t.Errorf("content is wrong: %s != %s", evt.Content, canonical.Content)
	}
	if evt.Kind != canonical.Kind {
		t.Errorf("kind is wrong: %d != %d", evt.Kind, canonical.Kind)
	}
	if evt.CreatedAt != Timestamp(canonical.CreatedAt.Unix()) {
		t.Errorf("created_at is wrong: %v != %v", evt.CreatedAt, canonical.CreatedAt)
	}
	if len(evt.Tags) != len(canonical.Tags) {
		t.Errorf("tag number is wrong: %v != %v", len(evt.Tags), len(canonical.Tags))
	}
	for i := range evt.Tags {
		if len(evt.Tags[i]) != len(canonical.Tags[i]) {
			t.Errorf("tag[%d] length is wrong: `%v` != `%v`", i, len(evt.Tags[i]), len(canonical.Tags[i]))
		}
		for j := range evt.Tags[i] {
			if evt.Tags[i][j] != canonical.Tags[i][j] {
				t.Errorf("tag[%d][%d] is wrong: `%s` != `%s`", i, j, evt.Tags[i][j], canonical.Tags[i][j])
			}
		}
	}
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

type tlvEvent struct {
	tlv   []byte
	event *Event
}

func loadEventsTLV() []tlvEvent {
	events := loadEvents()
	payloads := make([]tlvEvent, len(events))
	for i, evtstr := range events {
		evt := &Event{}
		json.Unmarshal([]byte(evtstr), evt)
		payloads[i].event = evt
		payloads[i].tlv = encodeEventTLV(evt)
	}
	return payloads
}

func loadEventsNson() []string {
	events := loadEvents()
	payloads := make([]string, len(events))
	for i, evtstr := range events {
		evt := &Event{}
		json.Unmarshal([]byte(evtstr), evt)
		payloads[i] = encodeNson(evt)
	}
	return payloads
}
