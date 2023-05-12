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
BenchmarkShortEvent/json.Unmarshal-4         	    4959	    569503 ns/op
BenchmarkShortEvent/gjson-4                  	   11474	    116819 ns/op
BenchmarkShortEvent/nson-4                   	   29349	     46547 ns/op
BenchmarkShortEvent/jsonparser-4             	    5716	    182513 ns/op
BenchmarkShortEvent/easyjson-4               	   10000	    118090 ns/op
BenchmarkShortEvent/ffjson-4                 	    6932	    144392 ns/op
BenchmarkShortEvent/go-json-4                	    5457	    199933 ns/op
BenchmarkShortEvent/ujson-4                  	    7508	    355112 ns/op
BenchmarkShortEvent/sonic-4                  	    5931	    219858 ns/op
BenchmarkShortEvent/sonic/get-4              	   32475	     48971 ns/op
BenchmarkShortEvent/sonic/searcher/get-4     	   25567	     48602 ns/op
BenchmarkShortEvent/simdjson-4               	     825	   1394563 ns/op
BenchmarkFullEvent/json.Unmarshal-4          	    3367	    502371 ns/op
BenchmarkFullEvent/gjson-4                   	    5132	    217851 ns/op
BenchmarkFullEvent/gjson_assign-4            	    3837	    275750 ns/op
BenchmarkFullEvent/jsonparser-4              	    2550	    409699 ns/op
BenchmarkFullEvent/jsonparser_assign-4       	    2017	    504660 ns/op
BenchmarkFullEvent/easyjson-4                	    9298	    115181 ns/op
BenchmarkFullEvent/ffjson-4                  	    5601	    193665 ns/op
BenchmarkFullEvent/go-json-4                 	    4227	    292374 ns/op
BenchmarkFullEvent/ujson-4                   	    2752	    402870 ns/op
BenchmarkFullEvent/sonic-4                   	    4657	    257283 ns/op
BenchmarkFullEvent/sonic/searcher/get-4      	    4412	    241117 ns/op
BenchmarkFullEvent/simdjson-4                	    1051	   1166769 ns/op
BenchmarkFullEvent/tlv_naïve-4               	    5668	    200668 ns/op
BenchmarkFullEvent/leaner_binary-4           	   54166	     22456 ns/op
BenchmarkFullEvent/leaner_to_string-4        	   21201	     59523 ns/op
BenchmarkFullEvent/nson-4                    	   18963	     60658 ns/op
BenchmarkEnvelope/json.Unmarshal-4           	    1075	   1052952 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-4      	    2980	    444283 ns/op
BenchmarkEnvelope/sonic-4                    	    4441	    319900 ns/op
BenchmarkEnvelope/easyjson-4                 	    2070	    601091 ns/op
BenchmarkEnvelope/gjson_+_easyjson-4         	    5572	    228761 ns/op
BenchmarkEnvelope/gjson_+_fastjson-4         	    1422	    786870 ns/op
BenchmarkEnvelope/gjson_+_sonic-4            	    4761	    363745 ns/op
BenchmarkEnvelope/gjson_+_nson-4             	    7230	    169663 ns/op
BenchmarkEnvelope/sonic_+_easyjson-4         	    5968	    179782 ns/op
BenchmarkEnvelope/sonic_+_nson-4             	    8408	    134826 ns/op
```

## custom decoders and encoders

this repository also contains the `tlv_naïve`, `leaner_binary` and `nson` implementations.

- `tlv_naïve` is just what happens when a beginner Go programmer tries to turn JSON into TLV in a somewhat-generic way.
- `leaner_binary` is the most optimized custom and yet very simple hand implementation I could do without doing any of the raw memory pointer access that probably cap'n'proto and others would.
- `nson` is a nice gimmick that gives a performance boost to normal JSON as long as the JSON creator is building the JSON object in a special way. The reader doesn't have to read it in a special way or be aware of the nsonic nature of the JSON blob it just got, though, therefore it is backwards-compatible.
