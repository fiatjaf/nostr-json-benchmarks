## latest results

```
goos: linux
goarch: amd64
pkg: github.com/fiatjaf/nostr-json-benchmarks
cpu: Intel(R) Core(TM) i5-2400 CPU @ 3.10GHz
BenchmarkFullEvent/json.Unmarshal-4                 5330            375625 ns/op
BenchmarkFullEvent/easyjson-4                      14581             79874 ns/op
BenchmarkFullEvent/go-json-4                       10000            160184 ns/op
BenchmarkFullEvent/sonic-4                         10000            138207 ns/op
BenchmarkEnvelope/json.Unmarshal-4                  1830            890396 ns/op
BenchmarkEnvelope/go-nostr-4                        7086            262845 ns/op
BenchmarkEnvelope/sonic-4                           6571            287370 ns/op
BenchmarkEnvelope/easyjson-4                        3535            561440 ns/op
BenchmarkEnvelope/gjson_+_easyjson-4                8037            240467 ns/op
BenchmarkEnvelope/gjson_+_sonic-4                   6625            274665 ns/op
BenchmarkEnvelope/sonic_+_easyjson-4                7909            234478 ns/op
PASS
```
