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
BenchmarkShortEvent/json.Unmarshal-4         	    4146	    532179 ns/op
BenchmarkShortEvent/gjson-4                  	   18818	     56823 ns/op
BenchmarkShortEvent/jsonparser-4             	   24924	     47408 ns/op
BenchmarkShortEvent/easyjson-4               	   18440	     70164 ns/op
BenchmarkShortEvent/ffjson-4                 	   16790	     80334 ns/op
BenchmarkShortEvent/go-json-4                	   10000	    152907 ns/op
BenchmarkShortEvent/ujson-4                  	    7736	    325845 ns/op
BenchmarkShortEvent/sonic-4                  	   10000	    180920 ns/op
BenchmarkShortEvent/sonic/get-4              	   40761	     43212 ns/op
BenchmarkShortEvent/sonic/searcher/get-4     	   36086	     41635 ns/op
BenchmarkShortEvent/simdjson-4               	    1088	    979647 ns/op
BenchmarkLazyEvent/gjson/0-4                 	   12975	     92240 ns/op
BenchmarkLazyEvent/gjson/1-4                 	   12165	    103735 ns/op
BenchmarkLazyEvent/gjson/2-4                 	    8865	    132323 ns/op
BenchmarkLazyEvent/gjson-megalazy/0-4        	   11661	    114954 ns/op
BenchmarkLazyEvent/gjson-megalazy/1-4        	   13657	    120260 ns/op
BenchmarkLazyEvent/gjson-megalazy/2-4        	   10000	    142147 ns/op
BenchmarkLazyEvent/sonic/0-4                 	   20316	     63973 ns/op
BenchmarkLazyEvent/sonic/1-4                 	   15451	     70049 ns/op
BenchmarkLazyEvent/sonic/2-4                 	   13402	     90294 ns/op
BenchmarkLazyEvent/sonic-megalazy/0-4        	   22298	     46759 ns/op
BenchmarkLazyEvent/sonic-megalazy/1-4        	    8941	    143444 ns/op
BenchmarkLazyEvent/sonic-megalazy/2-4        	    3590	    309322 ns/op
BenchmarkLazyEvent/sonic-not-lazy/0-4        	    4366	    253861 ns/op
BenchmarkLazyEvent/sonic-not-lazy/1-4        	    8530	    270917 ns/op
BenchmarkLazyEvent/sonic-not-lazy/2-4        	    4675	    273443 ns/op
BenchmarkFullEvent/json.Unmarshal-4          	    2088	    492002 ns/op
BenchmarkFullEvent/gjson-4                   	    5983	    193825 ns/op
BenchmarkFullEvent/gjson_assign-4            	    4618	    262272 ns/op
BenchmarkFullEvent/jsonparser-4              	    5623	    384807 ns/op
BenchmarkFullEvent/jsonparser_assign-4       	    2841	    446346 ns/op
BenchmarkFullEvent/easyjson-4                	    8985	    112913 ns/op
BenchmarkFullEvent/ffjson-4                  	    6670	    170878 ns/op
BenchmarkFullEvent/go-json-4                 	    7022	    264510 ns/op
BenchmarkFullEvent/ujson-4                   	    2944	    349426 ns/op
BenchmarkFullEvent/sonic-4                   	    8851	    201738 ns/op
BenchmarkFullEvent/sonic/searcher/get-4      	    9092	    147012 ns/op
BenchmarkFullEvent/simdjson-4                	     841	   1918874 ns/op
BenchmarkFullEvent/custom-4                  	   32329	     38135 ns/op
BenchmarkFullEvent/tlv_binary-4              	   28915	     43954 ns/op
BenchmarkEnvelope/json.Unmarshal-4           	    1290	    994963 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-4      	    3230	    471566 ns/op
BenchmarkEnvelope/sonic-4                    	    4486	    302790 ns/op
BenchmarkEnvelope/easyjson-4                 	    2331	    433049 ns/op
BenchmarkEnvelope/gjson_+_easyjson-4         	    9145	    138804 ns/op
BenchmarkEnvelope/gjson_+_fastjson-4         	    2953	    569067 ns/op
BenchmarkEnvelope/gjson_+_sonic-4            	    4146	    320105 ns/op
BenchmarkEnvelope/gjson_+_custom-4           	    6790	    173978 ns/op
```
