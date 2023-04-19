package benchmarks

import (
	"encoding/json"
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
			t.Errorf("tag[%d] length is wrong: %v != %v", i, len(evt.Tags[i]), len(canonical.Tags[i]))
		}
		for j := range evt.Tags[i] {
			if evt.Tags[i][j] != canonical.Tags[i][j] {
				t.Errorf("tag[%d][%d] is wrong: %s != %s", i, j, evt.Tags[i][j], canonical.Tags[i][j])
			}
		}
	}
}
