package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/mailru/easyjson"
)

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

func TestNson(t *testing.T) {
	for _, jevt := range loadEventsNson() {
		evt, err := decodeNson(jevt)
		if err != nil {
			t.Errorf("failed to parse '%s': %s", jevt, err)
		}

		// fmt.Println("\n\ncomparing:\n", jevt, "\n~\n", evt)
		checkParsedCorrectly(t, jevt, evt)
	}
}

func TestTLV(t *testing.T) {
	for _, te := range loadEventsTLV() {
		evt, err := decodeEventTLV(te.binary)
		if err != nil {
			t.Errorf("failed to parse '%v': %s", te.binary, err)
		}
		jevt, _ := json.Marshal(te.event)

		// fmt.Println("\n\ncomparing:\n", string(jevt), "\n~\n", evt)
		checkParsedCorrectly(t, string(jevt), evt)
	}
}

func TestLeanerBinary(t *testing.T) {
	for _, te := range loadEventsLeaner() {
		evt, err := leanerDecode(te.binary)
		if err != nil {
			t.Errorf("failed to parse '%v': %s", te.binary, err)
		}
		jevt, _ := json.Marshal(te.event)

		// fmt.Println("\n\ncomparing:\n", string(jevt), "\n~\n", evt)
		checkParsedCorrectly(t, string(jevt), evt.ToNormal())
	}
}
