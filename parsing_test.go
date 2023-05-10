package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/mailru/easyjson"
)

// this is currently broken::w

// func TestCustomParser(t *testing.T) {
// 	for _, jevt := range loadEvents() {
// 		evt, err := parseEvent(jevt)
// 		if err != nil {
// 			t.Errorf("failed to parse '%s': %s", jevt, err)
// 		}
//
// 		checkParsedCorrectly(t, jevt, evt)
// 	}
// }

func TestEasyJson(t *testing.T) {
	for _, jevt := range loadEvents() {
		var evt Event
		err := easyjson.Unmarshal([]byte(jevt), &evt)
		if err != nil {
			t.Errorf("failed to parse '%s': %s", jevt, err)
		}

		checkParsedCorrectly(t, jevt, &evt)
	}
}

func TestSonic(t *testing.T) {
	for _, jevt := range loadEvents() {
		var evt Event
		err := sonic.UnmarshalString(jevt, &evt)
		if err != nil {
			t.Errorf("failed to parse '%s': %s", jevt, err)
		}

		checkParsedCorrectly(t, jevt, &evt)
	}
}

func TestTLV(t *testing.T) {
	for _, te := range loadEventsTLV() {
		evt := decodeEventTLV(te.tlv)
		if evt == nil {
			t.Errorf("failed to parse '%v'", te.tlv)
		}
		jevt, _ := json.Marshal(te.event)
		checkParsedCorrectly(t, string(jevt), evt)
	}
}
