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

```
goos: darwin
goarch: arm64
pkg: github.com/fiatjaf/nostr-json-benchmarks
BenchmarkShortEvent/json.Unmarshal-8              4034     377482 ns/op
BenchmarkShortEvent/gjson-8                      21648      59256 ns/op
BenchmarkShortEvent/nson-8                       53888      19518 ns/op
BenchmarkShortEvent/jsonparser-8                 24476      54190 ns/op
BenchmarkShortEvent/easyjson-8                   22184      49096 ns/op
BenchmarkShortEvent/ffjson-8                     22018      58369 ns/op
BenchmarkShortEvent/go-json-8                     9308     127530 ns/op
BenchmarkShortEvent/ujson-8                       7048     197696 ns/op
BenchmarkShortEvent/sonic-8                       5396     269349 ns/op
BenchmarkShortEvent/sonic/get-8                  13868      77267 ns/op
BenchmarkShortEvent/sonic/searcher/get-8         15487      84715 ns/op
BenchmarkFullEvent/json.Unmarshal-8               4564     244337 ns/op
BenchmarkFullEvent/gjson-8                        8539     123492 ns/op
BenchmarkFullEvent/gjson_assign-8                10000     121072 ns/op
BenchmarkFullEvent/jsonparser-8                   7464     137948 ns/op
BenchmarkFullEvent/jsonparser_assign-8           10270     116280 ns/op
BenchmarkFullEvent/easyjson-8                    45086      26690 ns/op
BenchmarkFullEvent/ffjson-8                      28021      50835 ns/op
BenchmarkFullEvent/go-json-8                     13024     104233 ns/op
BenchmarkFullEvent/ujson-8                        9045     147797 ns/op
BenchmarkFullEvent/sonic-8                        4116     280009 ns/op
BenchmarkFullEvent/sonic/searcher/get-8           3549     357866 ns/op
BenchmarkFullEvent/tlv_naïve-8                    9016     157144 ns/op
BenchmarkFullEvent/leaner_binary-8               82825      13976 ns/op
BenchmarkFullEvent/leaner_to_string-8            42699      29387 ns/op
BenchmarkFullEvent/nson-8                        63736      23841 ns/op
BenchmarkEnvelope/json.Unmarshal-8                4018     435988 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-8           3996     252699 ns/op
BenchmarkEnvelope/sonic-8                         3135     382234 ns/op
BenchmarkEnvelope/easyjson-8                      6696     233853 ns/op
BenchmarkEnvelope/gjson_+_easyjson-8             12333     153952 ns/op
BenchmarkEnvelope/gjson_+_fastjson-8              1567     973647 ns/op
BenchmarkEnvelope/gjson_+_sonic-8                 1986     530190 ns/op
BenchmarkEnvelope/gjson_+_nson-8                 21490      78198 ns/op
BenchmarkEnvelope/sonic_+_easyjson-8             10614     124745 ns/op
BenchmarkEnvelope/sonic_+_nson-8                 10000     154750 ns/op
```

```
goos: darwin
goarch: arm64
pkg: github.com/fiatjaf/nostr-json-benchmarks
cpu: Macbook Pro 16" M2 Max
BenchmarkShortEvent/json.Unmarshal-12              10057            117773 ns/op
BenchmarkShortEvent/gjson-12                       45710             25890 ns/op
BenchmarkShortEvent/nson-12                       120819              9917 ns/op
BenchmarkShortEvent/jsonparser-12                  46545             25248 ns/op
BenchmarkShortEvent/easyjson-12                    47290             24822 ns/op
BenchmarkShortEvent/ffjson-12                      43297             28228 ns/op
BenchmarkShortEvent/go-json-12                     24748             48551 ns/op
BenchmarkShortEvent/ujson-12                       17973             62363 ns/op
BenchmarkShortEvent/sonic-12                        9894            119743 ns/op
BenchmarkShortEvent/sonic/get-12                   28562             41139 ns/op
BenchmarkShortEvent/sonic/searcher/get-12          29032             41373 ns/op
BenchmarkFullEvent/json.Unmarshal-12               10000            112748 ns/op
BenchmarkFullEvent/gjson-12                        16086             74164 ns/op
BenchmarkFullEvent/gjson_assign-12                 15632             76765 ns/op
BenchmarkFullEvent/jsonparser-12                   12982             92350 ns/op
BenchmarkFullEvent/jsonparser_assign-12            13189             90307 ns/op
BenchmarkFullEvent/easyjson-12                     58528             20100 ns/op
BenchmarkFullEvent/ffjson-12                       34897             34245 ns/op
BenchmarkFullEvent/go-json-12                      21699             55074 ns/op
BenchmarkFullEvent/ujson-12                        17041             69810 ns/op
BenchmarkFullEvent/sonic-12                         9025            127670 ns/op
BenchmarkFullEvent/sonic/searcher/get-12            9950            111432 ns/op
BenchmarkFullEvent/tlv_naïve-12                    31765             36481 ns/op
BenchmarkFullEvent/leaner_binary-12               255860              4068 ns/op
BenchmarkFullEvent/leaner_to_string-12            104157             11543 ns/op
BenchmarkFullEvent/nson-12                         81714             13854 ns/op
BenchmarkEnvelope/json.Unmarshal-12                 5241            220753 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-12           12205             96958 ns/op
BenchmarkEnvelope/sonic-12                          7530            152746 ns/op
BenchmarkEnvelope/easyjson-12                       9283            121171 ns/op
BenchmarkEnvelope/gjson_+_easyjson-12              31456             37932 ns/op
BenchmarkEnvelope/gjson_+_fastjson-12               7239            157654 ns/op
BenchmarkEnvelope/gjson_+_sonic-12                  7514            146805 ns/op
BenchmarkEnvelope/gjson_+_nson-12                  34935             31489 ns/op
BenchmarkEnvelope/sonic_+_easyjson-12              27753             42690 ns/op
BenchmarkEnvelope/sonic_+_nson-12                  32534             36169 ns/op
```

```
goos: darwin
goarch: amd64
pkg: github.com/fiatjaf/nostr-json-benchmarks
cpu: VirtualApple @ 2.50GHz
BenchmarkShortEvent/json.Unmarshal-8              6217     179714 ns/op
BenchmarkShortEvent/gjson-8                      36538      32226 ns/op
BenchmarkShortEvent/nson-8                       83272      14249 ns/op
BenchmarkShortEvent/jsonparser-8                 21170      55993 ns/op
BenchmarkShortEvent/easyjson-8                   31347      38016 ns/op
BenchmarkShortEvent/ffjson-8                     31155      38699 ns/op
BenchmarkShortEvent/go-json-8                    17224      69056 ns/op
BenchmarkShortEvent/ujson-8                      13654      87347 ns/op
BenchmarkShortEvent/sonic-8                      20325      58497 ns/op
BenchmarkShortEvent/sonic/get-8                  38076      31821 ns/op
BenchmarkShortEvent/sonic/searcher/get-8         38018      31307 ns/op
BenchmarkFullEvent/json.Unmarshal-8               7348     152086 ns/op
BenchmarkFullEvent/gjson-8                       13862      85955 ns/op
BenchmarkFullEvent/gjson_assign-8                13224      90134 ns/op
BenchmarkFullEvent/jsonparser-8                   6626     178087 ns/op
BenchmarkFullEvent/jsonparser_assign-8            7071     164603 ns/op
BenchmarkFullEvent/easyjson-8                    34314      34892 ns/op
BenchmarkFullEvent/ffjson-8                      19056      62083 ns/op
BenchmarkFullEvent/go-json-8                     13140      90951 ns/op
BenchmarkFullEvent/ujson-8                       12519      95873 ns/op
BenchmarkFullEvent/sonic-8                       14888      81750 ns/op
BenchmarkFullEvent/sonic/searcher/get-8          10000     102297 ns/op
BenchmarkFullEvent/tlv_naïve-8                   19000      63939 ns/op
BenchmarkFullEvent/leaner_binary-8              193753       6029 ns/op
BenchmarkFullEvent/leaner_to_string-8            73759      16142 ns/op
BenchmarkFullEvent/nson-8                        58666      20454 ns/op
BenchmarkEnvelope/json.Unmarshal-8                3736     316968 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-8           9044     132818 ns/op
BenchmarkEnvelope/sonic-8                        10000     111878 ns/op
BenchmarkEnvelope/easyjson-8                      6904     170993 ns/op
BenchmarkEnvelope/gjson_+_easyjson-8             20804      56676 ns/op
BenchmarkEnvelope/gjson_+_fastjson-8              4863     233191 ns/op
BenchmarkEnvelope/gjson_+_sonic-8                10000     103394 ns/op
BenchmarkEnvelope/gjson_+_nson-8                 29474      40307 ns/op
BenchmarkEnvelope/sonic_+_easyjson-8             18372      65401 ns/op
BenchmarkEnvelope/sonic_+_nson-8                 25036      48087 ns/op
```

## custom decoders and encoders

this repository also contains the `tlv_naïve`, `leaner_binary` and `nson` implementations.

- `tlv_naïve` is just what happens when a beginner Go programmer tries to turn JSON into TLV in a somewhat-generic way.
- `leaner_binary` is the most optimized custom and yet very simple hand implementation I could do without doing any of the raw memory pointer access that probably cap'n'proto and others would.
- `nson` is a nice gimmick that gives a performance boost to normal JSON as long as the JSON creator is building the JSON object in a special way. The reader doesn't have to read it in a special way or be aware of the nsonic nature of the JSON blob it just got, though, therefore it is backwards-compatible.

## javascript tests

the `javascript.js` file has a decoder for `leaner` and one for `nson`, which are benchmarked in comparison with `JSON.parse`. these are the results:

```
JSON.parse    1.84 µs/iter
decodeNson    1.98 µs/iter
leanerDecode  4.94 µs/iter
```

the only issue there is that the nson decoder is not doing string unescaping, which makes it slighly faster than it should be, but since it is already worse than `JSON.parse` without that I didn't bother to implement that.

this shows that for javascript apps it makes no sense to implement binary encoding -- but that NSON is fine as it can just be ignored at all times.
