package benchmarks

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/nbd-wtf/go-nostr"
	"golang.org/x/exp/rand"
)

func checkParsedCorrectly(t *testing.T, jevt string, evt *nostr.Event) (isBad bool) {
	var canonical nostr.Event
	json.Unmarshal([]byte(jevt), &canonical)

	if evt.ID != canonical.ID {
		t.Errorf("id is wrong: %s != %s", evt.ID, canonical.ID)
		isBad = true
	}
	if evt.PubKey != canonical.PubKey {
		t.Errorf("pubkey is wrong: %s != %s", evt.PubKey, canonical.PubKey)
		isBad = true
	}
	if evt.Sig != canonical.Sig {
		t.Errorf("sig is wrong: %s != %s", evt.Sig, canonical.Sig)
		isBad = true
	}
	if evt.Content != canonical.Content {
		t.Errorf("content is wrong: %s != %s", evt.Content, canonical.Content)
		isBad = true
	}
	if evt.Kind != canonical.Kind {
		t.Errorf("kind is wrong: %d != %d", evt.Kind, canonical.Kind)
		isBad = true
	}
	if evt.CreatedAt != canonical.CreatedAt {
		t.Errorf("created_at is wrong: %v != %v", evt.CreatedAt, canonical.CreatedAt)
		isBad = true
	}
	if len(evt.Tags) != len(canonical.Tags) {
		t.Errorf("tag number is wrong: %v != %v", len(evt.Tags), len(canonical.Tags))
		isBad = true
	}
	for i := range evt.Tags {
		if len(evt.Tags[i]) != len(canonical.Tags[i]) {
			t.Errorf("tag[%d] length is wrong: `%v` != `%v`", i, len(evt.Tags[i]), len(canonical.Tags[i]))
			isBad = true
		}
		for j := range evt.Tags[i] {
			if evt.Tags[i][j] != canonical.Tags[i][j] {
				t.Errorf("tag[%d][%d] is wrong: `%s` != `%s`", i, j, evt.Tags[i][j], canonical.Tags[i][j])
				isBad = true
			}
		}
	}

	return isBad
}

func loadLines() []string {
	b, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	text := string(b)
	lines := strings.Split(text, "\n")
	return lines[0 : len(lines)-1]
}

func loadEvents() [][]byte {
	lines := loadLines()
	events := make([][]byte, 0, len(lines))

	for _, line := range lines {
		if strings.HasPrefix(line, "[\"EVENT") {
			events = append(events, []byte(line[13:len(line)-1]))
		}
	}

	return events
}

func loadEnvelopes() [][]byte {
	events := loadEvents()
	envelopes := make([][]byte, len(events), len(events)*15/10)
	for i, evtstr := range events {
		envelopes[i] = []byte(fmt.Sprintf(`["EVENT","_",%s]`, evtstr))
	}
	for i := 0; i < len(events)/4; i++ {
		envelopes = append(envelopes, []byte(`["EOSE", "_kjasbd"]`))
	}
	for i := 0; i < len(events)/8; i++ {
		envelopes = append(envelopes, []byte(`["OK", "8c59239319637f97e007dad0d681e65ce35b1ace333b629e2d33f9465c132608", true]`))
	}
	for i := 0; i < len(events)/8; i++ {
		envelopes = append(envelopes, []byte(`["OK", "8c59239319637f97e007dad0d681e65ce35b1ace333b629e2d33f9465c132608", false, "askjdbkasjbdskajbd"]`))
	}
	for i := range envelopes {
		j := rand.Intn(i + 1)
		envelopes[i], envelopes[j] = envelopes[j], envelopes[i]
	}
	return envelopes
}
