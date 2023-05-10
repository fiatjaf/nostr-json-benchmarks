package benchmarks

import (
	"testing"
)

var nsonTestEvents = []string{
	`{"id":"ae1fc7154296569d87ca4663f6bdf448c217d1590d28c85d158557b8b43b4d69","pubkey":"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798","sig":"94e10947814b1ebe38af42300ecd90c7642763896c4f69506ae97bfdf54eec3c0c21df96b7d95daa74ff3d414b1d758ee95fc258125deebc31df0c6ba9396a51","created_at":1683660344,"nson":"2805000b0203000100400005040001004000000014","kind":30023,"content":"hello hello","tags":[["e","b6de44a9dd47d1c000f795ea0453046914f44ba7d5e369608b04867a575ea83e","reply"],["p","c26f7b252cea77a5b94f42b1a4771021be07d4df766407e47738605f7e3ab774","","wss://relay.damus.io"]]}`,
	`{"id":"ae1fc7154296569d87ca4663f6bdf448c217d1590d28c85d158557b8b43b4d69","pubkey":"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798","sig":"94e10947814b1ebe38af42300ecd90c7642763896c4f69506ae97bfdf54eec3c0c21df96b7d95daa74ff3d414b1d758ee95fc258125deebc31df0c6ba9396a51","created_at":1683660344,"nson":"280500100203000100400005040001004000000014","kind":30023,"content":"hello\n\"hello\"","tags":[["e","b6de44a9dd47d1c000f795ea0453046914f44ba7d5e369608b04867a575ea83e","reply"],["p","c26f7b252cea77a5b94f42b1a4771021be07d4df766407e47738605f7e3ab774","","wss://relay.damus.io"]]}`,
}

func TestBasicNsonParse(t *testing.T) {
	for _, jevt := range nsonTestEvents {
		evt, _ := decodeNson(jevt)
		checkParsedCorrectly(t, jevt, evt)
	}
}

func TestNsonPartialGet(t *testing.T) {
	for _, jevt := range nsonTestEvents {
		evt, _ := decodeNson(jevt)

		if id := nsonGetID(jevt); id != evt.ID {
			t.Error("partial id wrong")
		}
		if pubkey := nsonGetPubkey(jevt); pubkey != evt.PubKey {
			t.Error("partial pubkey wrong")
		}
		if sig := nsonGetSig(jevt); sig != evt.Sig {
			t.Error("partial sig wrong")
		}
		if createdAt := nsonGetCreatedAt(jevt); createdAt != evt.CreatedAt {
			t.Error("partial created_at wrong")
		}
		if kind := nsonGetKind(jevt); kind != evt.Kind {
			t.Error("partial kind wrong")
		}
		if content := nsonGetContent(jevt); content != evt.Content {
			t.Error("partial content wrong")
		}
	}
}
