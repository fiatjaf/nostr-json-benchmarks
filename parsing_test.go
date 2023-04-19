package benchmarks

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/mailru/easyjson"
)

var testEvents = []string{
	`{"content":"actually probably like $400 after labor","created_at":1681686919,"id":"0dcbf6b690c15057d9f2c23fbe2626f2b825621e72c518eaefe91ecc04c40011","kind":1,"pubkey":"a26d9180040977fbfa673374da58b057b2be33bd1b2615301969e2d7171a7c42","sig":"6926724001ad1f0392ac25e2caf609628fa79aa17e5c55d6c6050e71400560e2f47d736a283a43fc893b82124d7fe619e36851c738cd40c92a16f185a2ed2d34","tags":[["p","a26d9180040977fbfa673374da58b057b2be33bd1b2615301969e2d7171a7c42","wss://relay.mostr.pub"],["e","1d28e509c473609aec174156c635948096df5ca8d236e14b87f6b77050c9a6ea","wss://relay.mostr.pub","reply"],["mostr","https://poa.st/objects/a238e5fd-7575-4a20-bda3-144d84773ee3"]]}`,
	`{
    "content": "Let’s go heat!",
    "created_at": 1681689231,
    "id": "10f8463cfe8ba6a907a3500ca5c3ce6181f381802068997def6e0812c4beb64c",
    "kind": 1,
    "pubkey": "cc76679480a4504b963a3809cba60b458ebf068c62713621dda94b527860447d",
    "sig": "a02cccf9005ed528082c26f616989856e5ea8d9c8002b9b9d95f5a93f0bacefd90f5e64881622ba929694b3220daac646691c0ac0bb6c932c0ffb7eaeb0a902b",
    "tags": [
      [
        "e",
        "c24217b1318a2393eb4de4e7024396c33f4d1330bbbef3b07766068499abb490"
      ],
      [
        "p",
        "6475c0ba9c8f5f45dcfdae553189d1b8d089118295ba5b902c0a698e192f535b"
      ]
    ]
  }`,
	`{
    "content": "hello \"king of the jews\"",
    "created_at": 12,
    "id": "10f8463cfe8ba6a907a3500ca5c3ce6181f381802068997def6e0812c4beb64c",
    "kind": 30023,
    "pubkey": "cc76679480a4504b963a3809cba60b458ebf068c62713621dda94b527860447d",
    "sig": "a02cccf9005ed528082c26f616989856e5ea8d9c8002b9b9d95f5a93f0bacefd90f5e64881622ba929694b3220daac646691c0ac0bb6c932c0ffb7eaeb0a902b",
    "tags": []
  }`,
}

func TestCustomParser(t *testing.T) {
	for _, events := range [][]string{testEvents, loadEvents()} {
		for _, jevt := range events {
			evt, err := parseEvent(jevt)
			if err != nil {
				t.Errorf("failed to parse '%s': %s", jevt, err)
			}

			checkParsedCorrectly(t, jevt, evt)
		}
	}
}

func TestEasyJson(t *testing.T) {
	for _, events := range [][]string{testEvents, loadEvents()} {
		for _, jevt := range events {
			var evt Event
			err := easyjson.Unmarshal([]byte(jevt), &evt)
			if err != nil {
				t.Errorf("failed to parse '%s': %s", jevt, err)
			}

			checkParsedCorrectly(t, jevt, &evt)
		}
	}
}

func TestSonic(t *testing.T) {
	for _, events := range [][]string{testEvents, loadEvents()} {
		for _, jevt := range events {
			var evt Event
			err := sonic.UnmarshalString(jevt, &evt)
			if err != nil {
				t.Errorf("failed to parse '%s': %s", jevt, err)
			}

			checkParsedCorrectly(t, jevt, &evt)
		}
	}
}
