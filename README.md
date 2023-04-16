## latest results

```
goos: linux
goarch: amd64
pkg: github.com/fiatjaf/nostr-json-benchmarks
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkShortEvent/json.Unmarshal-8              7992     126756 ns/op
BenchmarkShortEvent/gjson-8                      59445      20438 ns/op
BenchmarkShortEvent/jsonparser-8                 55084      21448 ns/op
BenchmarkShortEvent/easyjson-8                   44335      26431 ns/op
BenchmarkShortEvent/ffjson-8                     39339      30645 ns/op
BenchmarkShortEvent/go-json-8                    20524      58168 ns/op
BenchmarkShortEvent/ujson-8                      15710      77279 ns/op
BenchmarkShortEvent/sonic-8                      25264      47235 ns/op
BenchmarkShortEvent/sonic/get-8                  97633      11802 ns/op
BenchmarkShortEvent/sonic/searcher/get-8        102620      11617 ns/op
BenchmarkLazyEvent/gjson/0-8                     50014      23433 ns/op
BenchmarkLazyEvent/gjson/1-8                     47962      25316 ns/op
BenchmarkLazyEvent/gjson/2-8                     36418      32128 ns/op
BenchmarkLazyEvent/gjson-megalazy/0-8            51171      23493 ns/op
BenchmarkLazyEvent/gjson-megalazy/1-8            47170      25012 ns/op
BenchmarkLazyEvent/gjson-megalazy/2-8            37058      32219 ns/op
BenchmarkLazyEvent/sonic/0-8                     76059      15751 ns/op
BenchmarkLazyEvent/sonic/1-8                     69402      17020 ns/op
BenchmarkLazyEvent/sonic/2-8                     54028      22084 ns/op
BenchmarkLazyEvent/sonic-megalazy/0-8            93540      12638 ns/op
BenchmarkLazyEvent/sonic-megalazy/1-8            29061      41409 ns/op
BenchmarkLazyEvent/sonic-megalazy/2-8            12225      97880 ns/op
BenchmarkLazyEvent/sonic-not-lazy/0-8            19933      59984 ns/op
BenchmarkLazyEvent/sonic-not-lazy/1-8            19168      67242 ns/op
BenchmarkLazyEvent/sonic-not-lazy/2-8            17490      68073 ns/op
BenchmarkFullEvent/json.Unmarshal-8               8437     134126 ns/op
BenchmarkFullEvent/gjson-8                       12558      91395 ns/op
BenchmarkFullEvent/gjson_assign-8                10000     101023 ns/op
BenchmarkFullEvent/jsonparser-8                  10000     115689 ns/op
BenchmarkFullEvent/jsonparser_assign-8            9474     119610 ns/op
BenchmarkFullEvent/easyjson-8                    42764      29632 ns/op
BenchmarkFullEvent/ffjson-8                      25148      46276 ns/op
BenchmarkFullEvent/go-json-8                     15232      79610 ns/op
BenchmarkFullEvent/ujson-8                       12438      95242 ns/op
BenchmarkFullEvent/sonic-8                       18808      63641 ns/op
BenchmarkGoNostrEventTyped/go-nostr_(fastjson)-8              7128     163147 ns/op
BenchmarkGoNostrEventTyped/sonic-8                           18529      63434 ns/op
BenchmarkGoNostrEventTyped/sonic/searcher/get-8              18110      66491 ns/op
BenchmarkGoNostrEventTyped/easyjson-8                        40179      29621 ns/op
BenchmarkEnvelope/json.Unmarshal-8                            4162     270675 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-8                       9735     118360 ns/op
BenchmarkEnvelope/sonic-8                                    14570      81534 ns/op
BenchmarkEnvelope/easyjson-8                                  7657     154056 ns/op
BenchmarkEnvelope/gjson_+_easyjson-8                         19718      59953 ns/op
BenchmarkEnvelope/gjson_+_fastjson-8                          6255     197929 ns/op
BenchmarkEnvelope/gjson_+_sonic-8                            12292      96553 ns/op
```

```
goos: darwin
goarch: amd64
pkg: github.com/fiatjaf/nostr-json-benchmarks
cpu: VirtualApple @ 2.50GHz
BenchmarkShortEvent/json.Unmarshal-8              6344     192395 ns/op
BenchmarkShortEvent/gjson-8                      52974      22386 ns/op
BenchmarkShortEvent/jsonparser-8                 33144      35594 ns/op
BenchmarkShortEvent/easyjson-8                   29486      40891 ns/op
BenchmarkShortEvent/ffjson-8                     31902      37149 ns/op
BenchmarkShortEvent/go-json-8                    17902      69153 ns/op
BenchmarkShortEvent/ujson-8                      13508      87936 ns/op
BenchmarkShortEvent/sonic-8                      20160      58223 ns/op
BenchmarkShortEvent/sonic/get-8                  44620      26023 ns/op
BenchmarkShortEvent/sonic/searcher/get-8         47250      25475 ns/op
BenchmarkLazyEvent/gjson/0-8                     42606      27756 ns/op
BenchmarkLazyEvent/gjson/1-8                     42033      27940 ns/op
BenchmarkLazyEvent/gjson/2-8                     33012      36033 ns/op
BenchmarkLazyEvent/gjson-megalazy/0-8            43329      27170 ns/op
BenchmarkLazyEvent/gjson-megalazy/1-8            42418      27887 ns/op
BenchmarkLazyEvent/gjson-megalazy/2-8            33705      35773 ns/op
BenchmarkLazyEvent/sonic/0-8                     38289      30556 ns/op
BenchmarkLazyEvent/sonic/1-8                     39391      30872 ns/op
BenchmarkLazyEvent/sonic/2-8                     29160      40955 ns/op
BenchmarkLazyEvent/sonic-megalazy/0-8            43404      27646 ns/op
BenchmarkLazyEvent/sonic-megalazy/1-8            13741      87149 ns/op
BenchmarkLazyEvent/sonic-megalazy/2-8             5533     213937 ns/op
BenchmarkLazyEvent/sonic-not-lazy/0-8            13936      86070 ns/op
BenchmarkLazyEvent/sonic-not-lazy/1-8            13908      86919 ns/op
BenchmarkLazyEvent/sonic-not-lazy/2-8            13684      88232 ns/op
BenchmarkFullEvent/json.Unmarshal-8               7237     163185 ns/op
BenchmarkFullEvent/gjson-8                       14703      80775 ns/op
BenchmarkFullEvent/gjson_assign-8                14001      85197 ns/op
BenchmarkFullEvent/jsonparser-8                   7188     167197 ns/op
BenchmarkFullEvent/jsonparser_assign-8            7284     162858 ns/op
BenchmarkFullEvent/easyjson-8                    32721      36832 ns/op
BenchmarkFullEvent/ffjson-8                      18822      63711 ns/op
BenchmarkFullEvent/go-json-8                     12494      94566 ns/op
BenchmarkFullEvent/ujson-8                       12168      97568 ns/op
BenchmarkFullEvent/sonic-8                       14277      83630 ns/op
BenchmarkGoNostrEventTyped/go-nostr_(fastjson)-8              5760     203313 ns/op
BenchmarkGoNostrEventTyped/sonic-8                           14433      84381 ns/op
BenchmarkGoNostrEventTyped/sonic/searcher/get-8              10000     110428 ns/op
BenchmarkGoNostrEventTyped/easyjson-8                        33031      36826 ns/op
BenchmarkEnvelope/json.Unmarshal-8                            3354     340486 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-8                       7682     147666 ns/op
BenchmarkEnvelope/sonic-8                                    10000     115144 ns/op
BenchmarkEnvelope/easyjson-8                                  6362     186066 ns/op
BenchmarkEnvelope/gjson_+_easyjson-8                         20208      59356 ns/op
BenchmarkEnvelope/gjson_+_fastjson-8                          5320     224552 ns/op
BenchmarkEnvelope/gjson_+_sonic-8                            10000     107003 ns/op
```
