## latest results

```
goos: linux
goarch: amd64
pkg: github.com/fiatjaf/nostr-json-benchmarks
cpu: AMD Ryzen 3 3200G with Radeon Vega Graphics
BenchmarkShortEvent/json.Unmarshal-4         	    5262	    494831 ns/op
BenchmarkShortEvent/gjson-4                  	   21432	     65843 ns/op
BenchmarkShortEvent/jsonparser-4             	   10000	    101185 ns/op
BenchmarkShortEvent/easyjson-4               	   10000	    107473 ns/op
BenchmarkShortEvent/ffjson-4                 	   10594	    138330 ns/op
BenchmarkShortEvent/go-json-4                	    4531	    237865 ns/op
BenchmarkShortEvent/ujson-4                  	    3238	    351381 ns/op
BenchmarkShortEvent/sonic-4                  	    5607	    198278 ns/op
BenchmarkShortEvent/sonic/get-4              	   26646	     42415 ns/op
BenchmarkShortEvent/sonic/searcher/get-4     	   35936	     38874 ns/op
BenchmarkShortEvent/simdjson-4               	    1051	   1183707 ns/op
BenchmarkLazyEvent/gjson/0-4                 	   10000	    106619 ns/op
BenchmarkLazyEvent/gjson/1-4                 	    9961	    113103 ns/op
BenchmarkLazyEvent/gjson/2-4                 	    7622	    148727 ns/op
BenchmarkLazyEvent/gjson-megalazy/0-4        	    9987	    111654 ns/op
BenchmarkLazyEvent/gjson-megalazy/1-4        	   12168	     97094 ns/op
BenchmarkLazyEvent/gjson-megalazy/2-4        	    7609	    138167 ns/op
BenchmarkLazyEvent/sonic/0-4                 	   20080	     62710 ns/op
BenchmarkLazyEvent/sonic/1-4                 	   16412	     68086 ns/op
BenchmarkLazyEvent/sonic/2-4                 	   13315	     91429 ns/op
BenchmarkLazyEvent/sonic-megalazy/0-4        	   25632	     46204 ns/op
BenchmarkLazyEvent/sonic-megalazy/1-4        	   10000	    157883 ns/op
BenchmarkLazyEvent/sonic-megalazy/2-4        	    3410	    388904 ns/op
BenchmarkLazyEvent/sonic-not-lazy/0-4        	    5121	    243446 ns/op
BenchmarkLazyEvent/sonic-not-lazy/1-4        	    4638	    234837 ns/op
BenchmarkLazyEvent/sonic-not-lazy/2-4        	    4431	    250808 ns/op
BenchmarkFullEvent/json.Unmarshal-4          	    2244	    462395 ns/op
BenchmarkFullEvent/gjson-4                   	    6865	    163783 ns/op
BenchmarkFullEvent/gjson_assign-4            	    5193	    199001 ns/op
BenchmarkFullEvent/jsonparser-4              	    2137	    508648 ns/op
BenchmarkFullEvent/jsonparser_assign-4       	    2530	    453301 ns/op
BenchmarkFullEvent/easyjson-4                	    8901	    116591 ns/op
BenchmarkFullEvent/ffjson-4                  	    6669	    175960 ns/op
BenchmarkFullEvent/go-json-4                 	    4050	    266537 ns/op
BenchmarkFullEvent/ujson-4                   	    2658	    413776 ns/op
BenchmarkFullEvent/sonic-4                   	    4890	    230371 ns/op
BenchmarkFullEvent/sonic/searcher/get-4      	    4299	    247482 ns/op
BenchmarkFullEvent/simdjson-4                	    1050	   1234070 ns/op
BenchmarkFullEvent/tlv_binary-4              	    6284	    197007 ns/op
BenchmarkFullEvent/nson-4                    	   15530	     80670 ns/op
BenchmarkEnvelope/json.Unmarshal-4           	    1284	    975399 ns/op
BenchmarkEnvelope/go-nostr_(fastjson)-4      	    2338	    445282 ns/op
BenchmarkEnvelope/sonic-4                    	    3733	    308238 ns/op
BenchmarkEnvelope/easyjson-4                 	    2550	    584880 ns/op
BenchmarkEnvelope/gjson_+_easyjson-4         	    4340	    237316 ns/op
BenchmarkEnvelope/gjson_+_fastjson-4         	    1647	    640919 ns/op
BenchmarkEnvelope/gjson_+_sonic-4            	    3126	    350412 ns/op
BenchmarkEnvelope/gjson_+_custom-4           	    7038	    176032 ns/op
BenchmarkEnvelope/sonic_+_easyjson-4         	    8539	    190145 ns/op
PASS
ok  	github.com/fiatjaf/nostr-json-benchmarks	73.120s
