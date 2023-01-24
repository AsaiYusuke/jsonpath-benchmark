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
Benchmark_threeLevelsWithIndex_oliveagle_JsonPath-2          	15649621	        76.94 ns/op	       0 B/op	       0 allocs/op
Benchmark_threeLevelsWithIndex_AsaiYusuke_JSONPath-2         	 7623028	       159.5 ns/op	      48 B/op	       3 allocs/op
Benchmark_threeLevelsWithIndex_ohler55_OjG_jp-2              	 2182900	       542.0 ns/op	    1040 B/op	       2 allocs/op
Benchmark_threeLevelsWithIndex_PaesslerAG_JSONPath-2         	 2127675	       580.6 ns/op	     208 B/op	       7 allocs/op
Benchmark_threeLevelsWithIndex_vmware_labs_YAML_JSONPath-2   	  885050	      1167 ns/op	     400 B/op	      25 allocs/op
Benchmark_threeLevelsWithIndex_bhmj_JSON_Slice-2             	  662049	      1813 ns/op	      24 B/op	       1 allocs/op
Benchmark_threeLevelsWithIndex_Spyzhov_Abstract_JSON-2       	  459765	      2466 ns/op	     759 B/op	      35 allocs/op
PASS
ok  	github.com/AsaiYusuke/jsonpath_benchmark	9.648s

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
Benchmark_recursiveDescentWithFilter_AsaiYusuke_JSONPath-2         	  430586	      2726 ns/op	     464 B/op	      17 allocs/op
Benchmark_recursiveDescentWithFilter_ohler55_OjG_jp-2              	  270266	      4362 ns/op	    5288 B/op	      21 allocs/op
Benchmark_recursiveDescentWithFilter_vmware_labs_YAML_JSONPath-2   	  181501	      6444 ns/op	    4528 B/op	     141 allocs/op
Benchmark_recursiveDescentWithFilter_bhmj_JSON_Slice-2             	   51892	     23056 ns/op	    1816 B/op	      43 allocs/op
Benchmark_recursiveDescentWithFilter_Spyzhov_Abstract_JSON-2       	   40300	     25391 ns/op	    7160 B/op	     279 allocs/op
PASS
ok  	github.com/AsaiYusuke/jsonpath_benchmark	6.431s

```
