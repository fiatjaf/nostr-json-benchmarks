## latest results

each "operation" in these benchmarks consists in decoding 35 events taken at a random time from random relays and stored at `data.json`.

- `BenchmarkShortEvent` is only decoding 3 fields from the event: `"created_at"`, `"content"` and `"pubkey"`.
- `BenchmarkFullEvent` is turning JSON into a struct (except on the case of `tlv_naïve` and `leaner_*` which take a binary-encoded format that isn't JSON).
- `BenchmarkEnvelope` is turning a relay message like `["EVENT", "_", {...}]` into a struct -- instead of just the event -- in this case multiple combinations of libraries are used sometimes, some for going through the array elements and others for actually decoding the event JSON.

```
goos: linux
goarch: amd64
pkg: github.com/fiatjaf/nostr-json-benchmarks
cpu: AMD Ryzen 3 3200G with Radeon Vega Graphics
BenchmarkShortEvent/json.Unmarshal-4                      	    4388	    535915 ns/op
BenchmarkShortEvent/gjson-4                               	   11971	    113192 ns/op
BenchmarkShortEvent/nson-4                                	   25970	     44606 ns/op
BenchmarkShortEvent/jsonparser-4                          	    5802	    175907 ns/op
BenchmarkShortEvent/easyjson-4                            	   10000	    111314 ns/op
BenchmarkShortEvent/ffjson-4                              	    7639	    144190 ns/op
BenchmarkShortEvent/go-json-4                             	    4581	    247255 ns/op
BenchmarkShortEvent/ujson-4                               	    3624	    403697 ns/op
BenchmarkShortEvent/sonic-4                               	    6135	    207432 ns/op
BenchmarkShortEvent/sonic/get-4                           	   24450	     43874 ns/op
BenchmarkShortEvent/sonic/searcher/get-4                  	   24254	     45064 ns/op
BenchmarkShortEvent/simdjson-4                            	     811	   1268456 ns/op
BenchmarkFullEvent/json.Unmarshal-4                       	    2103	    477904 ns/op
BenchmarkFullEvent/gjson-4                                	    5670	    203253 ns/op
BenchmarkFullEvent/gjson_assign-4                         	    4414	    238856 ns/op
BenchmarkFullEvent/jsonparser-4                           	    2648	    478627 ns/op
BenchmarkFullEvent/jsonparser_assign-4                    	    5835	    497328 ns/op
BenchmarkFullEvent/easyjson-4                             	    9217	    116945 ns/op
BenchmarkFullEvent/ffjson-4                               	    6764	    188987 ns/op
BenchmarkFullEvent/go-json-4                              	    4400	    276964 ns/op
BenchmarkFullEvent/ujson-4                                	    2638	    422623 ns/op
BenchmarkFullEvent/sonic-4                                	    4131	    253123 ns/op
BenchmarkFullEvent/sonic/searcher/get-4                   	    4363	    231676 ns/op
BenchmarkFullEvent/simdjson-4                             	    1057	   1179891 ns/op
BenchmarkFullEvent/tlv_naïve-4                            	    7887	    163227 ns/op
BenchmarkFullEvent/leaner_binary-4                        	   58695	     20018 ns/op
BenchmarkFullEvent/leaner_binary_to_stringified-4         	   20276	     49633 ns/op
BenchmarkFullEvent/nson-4                                 	   20380	     68420 ns/op
BenchmarkEnvelope/json.Unmarshal-4                        	    1348	    773058 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-4                   	    3045	    378101 ns/op
BenchmarkEnvelope/sonic-4                                 	    7434	    311607 ns/op
BenchmarkEnvelope/easyjson-4                              	    2365	    569178 ns/op
BenchmarkEnvelope/gjson_+_easyjson-4                      	    5320	    227517 ns/op
BenchmarkEnvelope/gjson_+_fastjson-4                      	    2222	    775710 ns/op
BenchmarkEnvelope/gjson_+_sonic-4                         	    3555	    367095 ns/op
BenchmarkEnvelope/gjson_+_nson-4                          	    8824	    180602 ns/op
BenchmarkEnvelope/sonic_+_easyjson-4                      	    5696	    200449 ns/op
BenchmarkEnvelope/sonic_+_nson-4                          	    8563	    151406 ns/op
```

## custom decoders and encoders

this repository also contains the `tlv_naïve`, `leaner_binary` and `nson` implementations.

- `tlv_naïve` is just what happens when a beginner Go programmer tries to turn JSON into TLV in a somewhat-generic way.
- `leaner_binary` is the most optimized custom and yet very simple hand implementation I could do without doing any of the raw memory pointer access that probably cap'n'proto and others would.
- `nson` is a nice gimmick that gives a performance boost to normal JSON as long as the JSON creator is building the JSON object in a special way. The reader doesn't have to read it in a special way or be aware of the nsonic nature of the JSON blob it just got, though, therefore it is backwards-compatible.
