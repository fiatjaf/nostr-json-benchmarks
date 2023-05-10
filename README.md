## latest results

```
goos: linux
goarch: amd64
pkg: github.com/fiatjaf/nostr-json-benchmarks
cpu: AMD Ryzen 3 3200G with Radeon Vega Graphics
BenchmarkShortEvent/json.Unmarshal-4         	    5284	    434455 ns/op
BenchmarkShortEvent/gjson-4                  	   16579	     77410 ns/op
BenchmarkShortEvent/nson-4                   	   45268	     28077 ns/op
BenchmarkShortEvent/jsonparser-4             	    9973	    125565 ns/op
BenchmarkShortEvent/easyjson-4               	   17586	     70463 ns/op
BenchmarkShortEvent/ffjson-4                 	   16285	     81429 ns/op
BenchmarkShortEvent/go-json-4                	   10000	    142931 ns/op
BenchmarkShortEvent/ujson-4                  	    7380	    202737 ns/op
BenchmarkShortEvent/sonic-4                  	    8863	    125568 ns/op
BenchmarkShortEvent/sonic/get-4              	   43250	     30467 ns/op
BenchmarkShortEvent/sonic/searcher/get-4     	   36811	     28518 ns/op
BenchmarkShortEvent/simdjson-4               	    1010	   1039704 ns/op
BenchmarkLazyEvent/gjson/0-4                 	   13604	    100188 ns/op
BenchmarkLazyEvent/gjson/1-4                 	   12079	    134587 ns/op
BenchmarkLazyEvent/gjson/2-4                 	   10051	    117537 ns/op
BenchmarkLazyEvent/gjson-megalazy/0-4        	   12824	     97576 ns/op
BenchmarkLazyEvent/gjson-megalazy/1-4        	   10000	    125462 ns/op
BenchmarkLazyEvent/gjson-megalazy/2-4        	   13950	     98287 ns/op
BenchmarkLazyEvent/sonic/0-4                 	   27982	     37991 ns/op
BenchmarkLazyEvent/sonic/1-4                 	   25592	     61873 ns/op
BenchmarkLazyEvent/sonic/2-4                 	   18328	     69789 ns/op
BenchmarkLazyEvent/sonic-megalazy/0-4        	   32838	     38287 ns/op
BenchmarkLazyEvent/sonic-megalazy/1-4        	   10000	    131118 ns/op
BenchmarkLazyEvent/sonic-megalazy/2-4        	    3392	    301577 ns/op
BenchmarkLazyEvent/sonic-not-lazy/0-4        	    6924	    217196 ns/op
BenchmarkLazyEvent/sonic-not-lazy/1-4        	    6468	    185756 ns/op
BenchmarkLazyEvent/sonic-not-lazy/2-4        	    6219	    205367 ns/op
BenchmarkFullEvent/json.Unmarshal-4          	    5547	    348039 ns/op
BenchmarkFullEvent/gjson-4                   	    6019	    177411 ns/op
BenchmarkFullEvent/gjson_assign-4            	    6432	    211247 ns/op
BenchmarkFullEvent/jsonparser-4              	    4455	    294369 ns/op
BenchmarkFullEvent/jsonparser_assign-4       	    4377	    256591 ns/op
BenchmarkFullEvent/easyjson-4                	   15625	     81747 ns/op
BenchmarkFullEvent/ffjson-4                  	   10000	    109502 ns/op
BenchmarkFullEvent/go-json-4                 	    8920	    173466 ns/op
BenchmarkFullEvent/ujson-4                   	    4820	    336796 ns/op
BenchmarkFullEvent/sonic-4                   	    7894	    163802 ns/op
BenchmarkFullEvent/sonic/searcher/get-4      	    7598	    199149 ns/op
BenchmarkFullEvent/simdjson-4                	    1135	   1044668 ns/op
BenchmarkFullEvent/tlv_binary-4              	    9349	    136897 ns/op
BenchmarkFullEvent/nson-4                    	   19155	     59467 ns/op
BenchmarkEnvelope/json.Unmarshal-4           	    2581	    546495 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-4      	    6091	    253763 ns/op
BenchmarkEnvelope/sonic-4                    	    8850	    202864 ns/op
BenchmarkEnvelope/easyjson-4                 	    3469	    375601 ns/op
BenchmarkEnvelope/gjson_+_easyjson-4         	    6013	    196086 ns/op
BenchmarkEnvelope/gjson_+_fastjson-4         	    1723	    649000 ns/op
BenchmarkEnvelope/gjson_+_sonic-4            	    4237	    325835 ns/op
BenchmarkEnvelope/gjson_+_nson-4             	    8834	    153968 ns/op
BenchmarkEnvelope/sonic_+_easyjson-4         	   11228	    141904 ns/op
BenchmarkEnvelope/sonic_+_nson-4             	   11680	    124066 ns/op
```
