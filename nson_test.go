package benchmarks

import (
	"testing"
)

var nsonTestEvents = []string{
	`{"id":"ae1fc7154296569d87ca4663f6bdf448c217d1590d28c85d158557b8b43b4d69","pubkey":"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798","sig":"94e10947814b1ebe38af42300ecd90c7642763896c4f69506ae97bfdf54eec3c0c21df96b7d95daa74ff3d414b1d758ee95fc258125deebc31df0c6ba9396a51","created_at":1683660344,"nson":"1a05000b02030140050401400014","kind":30023,"content":"hello hello","tags":[["e","b6de44a9dd47d1c000f795ea0453046914f44ba7d5e369608b04867a575ea83e","reply"],["p","c26f7b252cea77a5b94f42b1a4771021be07d4df766407e47738605f7e3ab774","","wss://relay.damus.io"]]}`,
}

func TestNsonParse(t *testing.T) {
	for _, jevt := range nsonTestEvents {
		evt := decodeNson(jevt)
		checkParsedCorrectly(t, jevt, evt)
	}
}
