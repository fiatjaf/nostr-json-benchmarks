## latest results

```
goos: darwin
goarch: arm64
pkg: github.com/fiatjaf/nostr-json-benchmarks
BenchmarkShortEvent/json.Unmarshal-10               9228            128891 ns/op
BenchmarkShortEvent/gjson-10                       61947             18971 ns/op
BenchmarkShortEvent/jsonparser-10                  66235             17718 ns/op
BenchmarkShortEvent/easyjson-10                    42387             28427 ns/op
BenchmarkShortEvent/ffjson-10                      39050             30530 ns/op
BenchmarkShortEvent/go-json-10                     22488             53260 ns/op
BenchmarkShortEvent/ujson-10                       17947             66888 ns/op
BenchmarkShortEvent/sonic-10                        8919            129733 ns/op
BenchmarkShortEvent/sonic/get-10                   44394             26252 ns/op
BenchmarkShortEvent/sonic/searcher/get-10          46618             25717 ns/op
BenchmarkShortEvent/simdjson-10                 1000000000               0.0000001 ns/op
BenchmarkLazyEvent/gjson/0-10                      51361             23441 ns/op
BenchmarkLazyEvent/gjson/1-10                      52131             23005 ns/op
BenchmarkLazyEvent/gjson/2-10                      38524             35013 ns/op
BenchmarkLazyEvent/gjson-megalazy/0-10             51414             23362 ns/op
BenchmarkLazyEvent/gjson-megalazy/1-10             51658             22900 ns/op
BenchmarkLazyEvent/gjson-megalazy/2-10             39837             30046 ns/op
BenchmarkLazyEvent/sonic/0-10                      38512             30966 ns/op
BenchmarkLazyEvent/sonic/1-10                      40417             29519 ns/op
BenchmarkLazyEvent/sonic/2-10                      30669             39383 ns/op
BenchmarkLazyEvent/sonic-megalazy/0-10             40968             29435 ns/op
BenchmarkLazyEvent/sonic-megalazy/1-10             13112             91279 ns/op
BenchmarkLazyEvent/sonic-megalazy/2-10              5338            223941 ns/op
BenchmarkLazyEvent/sonic-not-lazy/0-10              8670            137567 ns/op
BenchmarkLazyEvent/sonic-not-lazy/1-10              8498            138604 ns/op
BenchmarkLazyEvent/sonic-not-lazy/2-10              8762            139614 ns/op
BenchmarkFullEvent/json.Unmarshal-10                9808            121191 ns/op
BenchmarkFullEvent/gjson-10                        15264             78048 ns/op
BenchmarkFullEvent/gjson_assign-10                 14671             80852 ns/op
BenchmarkFullEvent/jsonparser-10                   10000            101200 ns/op
BenchmarkFullEvent/jsonparser_assign-10            10000            102325 ns/op
BenchmarkFullEvent/easyjson-10                     56974             20769 ns/op
BenchmarkFullEvent/ffjson-10                       31890             37393 ns/op
BenchmarkFullEvent/go-json-10                      19761             60730 ns/op
BenchmarkFullEvent/ujson-10                        16048             74830 ns/op
BenchmarkFullEvent/sonic-10                         8191            136383 ns/op
BenchmarkFullEvent/sonic/searcher/get-10           10000            106450 ns/op
BenchmarkFullEvent/simdjson-10                  1000000000               0.0000000 ns/op
BenchmarkFullEvent/custom-10                       52083             23110 ns/op
BenchmarkFullEvent/tlv_binary-10                  114582             10824 ns/op
BenchmarkEnvelope/json.Unmarshal-10                 4732            239134 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-10           10000            106707 ns/op
BenchmarkEnvelope/sonic-10                          6950            162521 ns/op
BenchmarkEnvelope/easyjson-10                       8516            133234 ns/op
BenchmarkEnvelope/gjson_+_easyjson-10              24489             45128 ns/op
BenchmarkEnvelope/gjson_+_fastjson-10               7131            164180 ns/op
BenchmarkEnvelope/gjson_+_sonic-10                  7208            169623 ns/op
BenchmarkEnvelope/gjson_+_custom-10                25899             45427 ns/op
```

```
goos: linux
goarch: amd64
pkg: github.com/fiatjaf/nostr-json-benchmarks
cpu: AMD Ryzen 3 3200G with Radeon Vega Graphics
BenchmarkShortEvent/json.Unmarshal-4         	    3147	    589632 ns/op
BenchmarkShortEvent/gjson-4                  	   21421	     70428 ns/op
BenchmarkShortEvent/jsonparser-4             	   10000	    102389 ns/op
BenchmarkShortEvent/easyjson-4               	    8562	    125502 ns/op
BenchmarkShortEvent/ffjson-4                 	    7982	    145741 ns/op
BenchmarkShortEvent/go-json-4                	    4488	    247062 ns/op
BenchmarkShortEvent/ujson-4                  	    3614	    367802 ns/op
BenchmarkShortEvent/sonic-4                  	    7117	    215547 ns/op
BenchmarkShortEvent/sonic/get-4              	   31156	     42600 ns/op
BenchmarkShortEvent/sonic/searcher/get-4     	   34498	     44854 ns/op
BenchmarkShortEvent/simdjson-4               	    1156	   1225788 ns/op
BenchmarkLazyEvent/gjson/0-4                 	    8983	    117474 ns/op
BenchmarkLazyEvent/gjson/1-4                 	   10872	    114302 ns/op
BenchmarkLazyEvent/gjson/2-4                 	   10000	    144905 ns/op
BenchmarkLazyEvent/gjson-megalazy/0-4        	   10822	     99571 ns/op
BenchmarkLazyEvent/gjson-megalazy/1-4        	   10000	    103567 ns/op
BenchmarkLazyEvent/gjson-megalazy/2-4        	    7260	    148636 ns/op
BenchmarkLazyEvent/sonic/0-4                 	   22981	     61426 ns/op
BenchmarkLazyEvent/sonic/1-4                 	   16954	     76296 ns/op
BenchmarkLazyEvent/sonic/2-4                 	   13980	     92788 ns/op
BenchmarkLazyEvent/sonic-megalazy/0-4        	   24402	     47966 ns/op
BenchmarkLazyEvent/sonic-megalazy/1-4        	   10000	    149327 ns/op
BenchmarkLazyEvent/sonic-megalazy/2-4        	    7554	    332413 ns/op
BenchmarkLazyEvent/sonic-not-lazy/0-4        	    5456	    268631 ns/op
BenchmarkLazyEvent/sonic-not-lazy/1-4        	    4020	    281939 ns/op
BenchmarkLazyEvent/sonic-not-lazy/2-4        	    5419	    267253 ns/op
BenchmarkFullEvent/json.Unmarshal-4          	    2403	    494438 ns/op
BenchmarkFullEvent/gjson-4                   	    5269	    192318 ns/op
BenchmarkFullEvent/gjson_assign-4            	    7635	    244467 ns/op
BenchmarkFullEvent/jsonparser-4              	    2998	    397891 ns/op
BenchmarkFullEvent/jsonparser_assign-4       	    4479	    398433 ns/op
BenchmarkFullEvent/easyjson-4                	   10000	    105923 ns/op
BenchmarkFullEvent/ffjson-4                  	    7736	    158215 ns/op
BenchmarkFullEvent/go-json-4                 	    6388	    199642 ns/op
BenchmarkFullEvent/ujson-4                   	    2763	    382812 ns/op
BenchmarkFullEvent/sonic-4                   	    5007	    251413 ns/op
BenchmarkFullEvent/sonic/searcher/get-4      	    5062	    217612 ns/op
BenchmarkFullEvent/simdjson-4                	    1164	    986608 ns/op
BenchmarkFullEvent/custom-4                  	   23364	     52946 ns/op
BenchmarkFullEvent/tlv_binary-4              	   33868	     39878 ns/op
BenchmarkEnvelope/json.Unmarshal-4           	    1665	   1086680 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-4      	    2487	    478782 ns/op
BenchmarkEnvelope/sonic-4                    	    8628	    329951 ns/op
BenchmarkEnvelope/easyjson-4                 	    4176	    611381 ns/op
BenchmarkEnvelope/gjson_+_easyjson-4         	    6109	    215240 ns/op
BenchmarkEnvelope/gjson_+_fastjson-4         	    1920	    697164 ns/op
BenchmarkEnvelope/gjson_+_sonic-4            	    6285	    255959 ns/op
BenchmarkEnvelope/gjson_+_custom-4           	   10000	    134707 ns/op
BenchmarkEnvelope/sonic_+_easyjson-4         	   10000	    158167 ns/op
```
