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
cpu: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz
Benchmark1_oliveagle_JsonPath-2          	15344602	        78.41 ns/op	       0 B/op	       0 allocs/op
Benchmark1_AsaiYusuke_JSONPath-2         	 8699992	       139.9 ns/op	      24 B/op	       2 allocs/op
Benchmark1_ohler55_OjG_jp-2              	 2162139	       568.0 ns/op	    1040 B/op	       2 allocs/op
Benchmark1_PaesslerAG_JSONPath-2         	 2119836	       606.3 ns/op	     208 B/op	       7 allocs/op
Benchmark1_vmware_labs_YAML_JSONPath-2   	  895326	      1179 ns/op	     400 B/op	      25 allocs/op
Benchmark1_bhmj_JSON_Slice-2             	  659169	      1817 ns/op	      24 B/op	       1 allocs/op
Benchmark1_Spyzhov_Abstract_JSON-2       	  452569	      2479 ns/op	     759 B/op	      35 allocs/op
PASS
ok  	github.com/AsaiYusuke/jsonpath_benchmark	9.745s

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
cpu: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz
Benchmark2_AsaiYusuke_JSONPath-2         	  497533	      2386 ns/op	     240 B/op	       9 allocs/op
Benchmark2_ohler55_OjG_jp-2              	  270592	      4385 ns/op	    5288 B/op	      21 allocs/op
Benchmark2_vmware_labs_YAML_JSONPath-2   	  179518	      6465 ns/op	    4528 B/op	     141 allocs/op
Benchmark2_bhmj_JSON_Slice-2             	   51969	     22948 ns/op	    1816 B/op	      43 allocs/op
Benchmark2_Spyzhov_Abstract_JSON-2       	   47300	     26757 ns/op	    7160 B/op	     279 allocs/op
PASS
ok  	github.com/AsaiYusuke/jsonpath_benchmark	6.644s

```
