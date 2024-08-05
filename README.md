# JSONPath Benchmark

[![Benchmark JSONPath](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml/badge.svg)](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml)

I benchmarked two JSONPaths using several libraries for the Go language.
Results can be viewed through Github actions.

The focus of the measurement is the cost per job for a task that involves a significant amount of looping after all the initial preparations are complete.

There were notable performance differences.
However, if the number of queries is limited, there will not be a significant variation between any of them.

Additionally, the results may vary depending on the input data.
Therefore, this benchmark should be considered for informational purposes only and re-evaluated at regular intervals.

- [AsaiYusuke/JSONPath](https://github.com/AsaiYusuke/jsonpath)
- [ohler55/OjG](https://github.com/ohler55/ojg)
- [vmware-labs/YAML JSONPath](https://github.com/vmware-labs/yaml-jsonpath)
- [bhmj/JSONSlice](https://github.com/bhmj/jsonslice)
- [Spyzhov/Abstract JSON](https://github.com/spyzhov/ajson)
- [oliveagle/JsonPath](https://github.com/oliveagle/jsonpath)
- [PaesslerAG/JSONPath](https://github.com/PaesslerAG/jsonpath)

## Comparing across all libraries

```
JSONPath : $.store.book[0].price
```

The following is the outcome of a JSONPath that was processed by all libraries.
The library "oliveagle/JsonPath" performed the fastest, while my own library placed second.

```
goos: linux
goarch: amd64
pkg: github.com/AsaiYusuke/jsonpath_benchmark
cpu: AMD EPYC 7763 64-Core Processor                
Benchmark1_oliveagle_JsonPath-4          	19315380	        62.11 ns/op	       0 B/op	       0 allocs/op
Benchmark1_AsaiYusuke_JSONPath-4         	 9839335	       121.5 ns/op	      24 B/op	       2 allocs/op
Benchmark1_ohler55_OjG_jp-4              	 2994937	       395.1 ns/op	    1040 B/op	       2 allocs/op
Benchmark1_PaesslerAG_JSONPath-4         	 2732386	       439.1 ns/op	     208 B/op	       7 allocs/op
Benchmark1_vmware_labs_YAML_JSONPath-4   	 1362562	       887.0 ns/op	     400 B/op	      25 allocs/op
Benchmark1_bhmj_JSON_Slice-4             	  819874	      1432 ns/op	      24 B/op	       1 allocs/op
Benchmark1_Spyzhov_Abstract_JSON-4       	  583876	      1839 ns/op	     760 B/op	      35 allocs/op
PASS
ok  	github.com/AsaiYusuke/jsonpath_benchmark	10.216s

```

## Comparing with a slightly complex syntax

```
JSONPath : $..book[?(@.price > $.store.bicycle.price)]
```

Among the limited number of libraries capable of handling complex syntax, my library performed the fastest.

```
goos: linux
goarch: amd64
pkg: github.com/AsaiYusuke/jsonpath_benchmark
cpu: AMD EPYC 7763 64-Core Processor                
Benchmark2_AsaiYusuke_JSONPath-4         	  656502	      1792 ns/op	     240 B/op	       9 allocs/op
Benchmark2_ohler55_OjG_jp-4              	  295694	      3897 ns/op	    5368 B/op	      25 allocs/op
Benchmark2_vmware_labs_YAML_JSONPath-4   	  230881	      4952 ns/op	    4528 B/op	     141 allocs/op
Benchmark2_bhmj_JSON_Slice-4             	   70855	     16750 ns/op	    1816 B/op	      43 allocs/op
Benchmark2_Spyzhov_Abstract_JSON-4       	   61496	     19286 ns/op	    7160 B/op	     279 allocs/op
PASS
ok  	github.com/AsaiYusuke/jsonpath_benchmark	6.339s

```
