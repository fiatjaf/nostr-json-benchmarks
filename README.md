```
goos: linux
goarch: amd64
pkg: github.com/fiatjaf/nostr-json-benchmarks
cpu: Intel(R) Core(TM) i5-2400 CPU @ 3.10GHz
BenchmarkShortEvent/json.Unmarshal-4         	    2426	    503515 ns/op
BenchmarkShortEvent/gjson-4                  	   15012	     82522 ns/op
BenchmarkShortEvent/jsonparser-4             	   13888	     77707 ns/op
BenchmarkShortEvent/easyjson-4               	   10000	    106886 ns/op
BenchmarkShortEvent/ffjson-4                 	    9321	    124589 ns/op
BenchmarkShortEvent/go-json-4                	    8940	    225092 ns/op
BenchmarkShortEvent/ujson-4                  	    6625	    300135 ns/op
BenchmarkShortEvent/sonic-4                  	   10000	    164359 ns/op
BenchmarkShortEvent/sonic/get-4              	   28290	     39486 ns/op
BenchmarkShortEvent/sonic/searcher/get-4     	   29724	     37390 ns/op
BenchmarkLazyEvent/gjson/0-4                 	    8590	    132401 ns/op
BenchmarkLazyEvent/gjson/1-4                 	   12025	    116181 ns/op
BenchmarkLazyEvent/gjson/2-4                 	    6637	    164679 ns/op
BenchmarkLazyEvent/gjson-megalazy/0-4        	   11526	    129139 ns/op
BenchmarkLazyEvent/gjson-megalazy/1-4        	   11755	    109693 ns/op
BenchmarkLazyEvent/gjson-megalazy/2-4        	    7406	    146907 ns/op
BenchmarkLazyEvent/sonic/0-4                 	   23642	     57547 ns/op
BenchmarkLazyEvent/sonic/1-4                 	   17113	     67337 ns/op
BenchmarkLazyEvent/sonic/2-4                 	   16078	     78742 ns/op
BenchmarkLazyEvent/sonic-megalazy/0-4        	   23032	     45484 ns/op
BenchmarkLazyEvent/sonic-megalazy/1-4        	    7257	    162680 ns/op
BenchmarkLazyEvent/sonic-megalazy/2-4        	    4239	    335758 ns/op
BenchmarkLazyEvent/sonic-not-lazy/0-4        	    7305	    227692 ns/op
BenchmarkLazyEvent/sonic-not-lazy/1-4        	    7372	    204597 ns/op
BenchmarkLazyEvent/sonic-not-lazy/2-4        	    4500	    233741 ns/op
BenchmarkFullEvent/json.Unmarshal-4          	    2838	    473989 ns/op
BenchmarkFullEvent/gjson-4                   	    4735	    224242 ns/op
BenchmarkFullEvent/gjson_assign-4            	    5478	    281850 ns/op
BenchmarkFullEvent/jsonparser-4              	    3355	    424668 ns/op
BenchmarkFullEvent/jsonparser_assign-4       	    3067	    386219 ns/op
BenchmarkFullEvent/easyjson-4                	   11836	     92130 ns/op
BenchmarkFullEvent/ffjson-4                  	   10000	    143380 ns/op
BenchmarkFullEvent/go-json-4                 	    3924	    270988 ns/op
BenchmarkFullEvent/ujson-4                   	    2928	    350827 ns/op
BenchmarkFullEvent/sonic-4                   	    8581	    221211 ns/op
BenchmarkGoNostrEventTyped/go-nostr_(fastjson)-4         	    2232	    624558 ns/op
BenchmarkGoNostrEventTyped/sonic-4                       	    5652	    243128 ns/op
BenchmarkGoNostrEventTyped/sonic/searcher/get-4          	    5983	    205589 ns/op
BenchmarkGoNostrEventTyped/easyjson-4                    	   12874	     94925 ns/op
BenchmarkEnvelope/json.Unmarshal-4                       	    1350	   1001001 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-4                  	    4542	    380435 ns/op
BenchmarkEnvelope/sonic-4                                	    6600	    300537 ns/op
BenchmarkEnvelope/easyjson-4                             	    2294	    476680 ns/op
BenchmarkEnvelope/gjson_+_easyjson-4                     	    4464	    240154 ns/op
BenchmarkEnvelope/gjson_+_fastjson-4                     	    2588	    723539 ns/op
BenchmarkEnvelope/gjson_+_sonic-4                        	    3758	    344509 ns/op
```

```
goos: linux
goarch: amd64
pkg: github.com/fiatjaf/nostr-json-benchmarks
cpu: AMD Ryzen 3 3200G with Radeon Vega Graphics
BenchmarkShortEvent/json.Unmarshal-4         	    4198	    543103 ns/op
BenchmarkShortEvent/gjson-4                  	   15984	     73074 ns/op
BenchmarkShortEvent/jsonparser-4             	   12283	    106477 ns/op
BenchmarkShortEvent/easyjson-4               	    9250	    126239 ns/op
BenchmarkShortEvent/ffjson-4                 	    8499	    144808 ns/op
BenchmarkShortEvent/go-json-4                	    5812	    252707 ns/op
BenchmarkShortEvent/ujson-4                  	    2725	    377234 ns/op
BenchmarkShortEvent/sonic-4                  	    4632	    229008 ns/op
BenchmarkShortEvent/sonic/get-4              	   46164	     38560 ns/op
BenchmarkShortEvent/sonic/searcher/get-4     	   28123	     40713 ns/op
BenchmarkLazyEvent/gjson/0-4                 	   10945	    103520 ns/op
BenchmarkLazyEvent/gjson/1-4                 	    9576	    110570 ns/op
BenchmarkLazyEvent/gjson/2-4                 	    7506	    133474 ns/op
BenchmarkLazyEvent/gjson-megalazy/0-4        	   12618	    110834 ns/op
BenchmarkLazyEvent/gjson-megalazy/1-4        	   11730	     98559 ns/op
BenchmarkLazyEvent/gjson-megalazy/2-4        	    8503	    126718 ns/op
BenchmarkLazyEvent/sonic/0-4                 	   18250	     66081 ns/op
BenchmarkLazyEvent/sonic/1-4                 	   14512	     79702 ns/op
BenchmarkLazyEvent/sonic/2-4                 	   13284	     85962 ns/op
BenchmarkLazyEvent/sonic-megalazy/0-4        	   22660	     54825 ns/op
BenchmarkLazyEvent/sonic-megalazy/1-4        	    9093	    162800 ns/op
BenchmarkLazyEvent/sonic-megalazy/2-4        	    4075	    310729 ns/op
BenchmarkLazyEvent/sonic-not-lazy/0-4        	    3710	    274257 ns/op
BenchmarkLazyEvent/sonic-not-lazy/1-4        	    3883	    274258 ns/op
BenchmarkLazyEvent/sonic-not-lazy/2-4        	    4029	    262577 ns/op
BenchmarkFullEvent/json.Unmarshal-4          	    2588	    503170 ns/op
BenchmarkFullEvent/gjson-4                   	    5896	    206969 ns/op
BenchmarkFullEvent/gjson_assign-4            	    4537	    231934 ns/op
BenchmarkFullEvent/jsonparser-4              	    5643	    338932 ns/op
BenchmarkFullEvent/jsonparser_assign-4       	    3063	    452436 ns/op
BenchmarkFullEvent/easyjson-4                	    8450	    124422 ns/op
BenchmarkFullEvent/ffjson-4                  	    5846	    194116 ns/op
BenchmarkFullEvent/go-json-4                 	    3598	    300389 ns/op
BenchmarkFullEvent/ujson-4                   	    2582	    401295 ns/op
BenchmarkFullEvent/sonic-4                   	    4593	    266281 ns/op
BenchmarkGoNostrEventTyped/go-nostr_(fastjson)-4         	    1968	    646721 ns/op
BenchmarkGoNostrEventTyped/sonic-4                       	    6648	    262897 ns/op
BenchmarkGoNostrEventTyped/sonic/searcher/get-4          	    4214	    287286 ns/op
BenchmarkGoNostrEventTyped/easyjson-4                    	    9390	    122513 ns/op
BenchmarkEnvelope/json.Unmarshal-4                       	    1185	    993432 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-4                  	    2926	    427638 ns/op
BenchmarkEnvelope/sonic-4                                	    3366	    356463 ns/op
BenchmarkEnvelope/easyjson-4                             	    1725	    653290 ns/op
BenchmarkEnvelope/gjson_+_easyjson-4                     	    4626	    245818 ns/op
BenchmarkEnvelope/gjson_+_fastjson-4                     	    1549	    750301 ns/op
BenchmarkEnvelope/gjson_+_sonic-4                        	    3087	    373509 ns/op
```
