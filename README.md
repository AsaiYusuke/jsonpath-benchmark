# JSONPath Benchmark

[![Benchmark JSONPath](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml/badge.svg)](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml)

This project benchmarks multiple JSONPath libraries written in Go.
It evaluates performance in specific scenarios and offers guidance for selecting an appropriate library.

## Overview

Results are generated via GitHub Actions for consistency and automation.

The benchmark focuses on the per-operation cost after initial setup is complete.
Preparatory steps, such as parsing or preloading data, are excluded from the measurements.
Instead, it emphasizes scenarios with intensive looping or query execution within the main operation.

Results may vary with factors such as input structure, query complexity, and runtime environment.
Treat these benchmarks as a general reference and re-evaluate them periodically.

## Libraries Benchmarked

The following libraries are included in this benchmark:

- [AsaiYusuke/JSONPath](https://github.com/AsaiYusuke/jsonpath)
- [ohler55/OjG](https://github.com/ohler55/ojg)
- [vmware-labs/YAML JSONPath](https://github.com/vmware-labs/yaml-jsonpath)
- [bhmj/JSONSlice](https://github.com/bhmj/jsonslice)
- [Spyzhov/Abstract JSON](https://github.com/spyzhov/ajson)
- [oliveagle/JsonPath](https://github.com/oliveagle/jsonpath)
- [PaesslerAG/JSONPath](https://github.com/PaesslerAG/jsonpath)

## Results: Simple Query

JSONPath:

``` text
$.store.book[0].price
```

Performance Summary:

All libraries support this query, enabling a direct performance comparison across them.
With buffer reuse enabled, AsaiYusuke/JSONPath is the fastest; when allocating a new buffer per operation, it ranks second.

``` bash
goos: linux
goarch: amd64
pkg: github.com/AsaiYusuke/jsonpath_benchmark
cpu: AMD EPYC 7763 64-Core Processor                
Benchmark1_AsaiYusuke_JSONPath_reuseBuffer-4   	18992085	        62.50 ns/op	       0 B/op	       0 allocs/op
Benchmark1_oliveagle_JsonPath-4                	16875874	        70.56 ns/op	       0 B/op	       0 allocs/op
Benchmark1_AsaiYusuke_JSONPath-4               	11419996	       103.9 ns/op	      16 B/op	       1 allocs/op
Benchmark1_ohler55_OjG_jp-4                    	 3395434	       350.9 ns/op	    1168 B/op	       2 allocs/op
Benchmark1_PaesslerAG_JSONPath-4               	 2958105	       399.7 ns/op	     208 B/op	       7 allocs/op
Benchmark1_vmware_labs_YAML_JSONPath-4         	 1323622	       907.0 ns/op	     464 B/op	      28 allocs/op
Benchmark1_bhmj_JSON_Slice-4                   	  929371	      1286 ns/op	      24 B/op	       1 allocs/op
Benchmark1_Spyzhov_Abstract_JSON-4             	  863251	      1431 ns/op	     472 B/op	      25 allocs/op
PASS
ok  	github.com/AsaiYusuke/jsonpath_benchmark	9.578s

```

## Results: Complex Query

JSONPath:

``` text
$..book[?(@.price > $.store.bicycle.price)]
```

Performance Summary:

This query uses more complex syntax, and only a subset of libraries could process it.
Among them, AsaiYusuke/JSONPath delivered the best performance.

``` bash
goos: linux
goarch: amd64
pkg: github.com/AsaiYusuke/jsonpath_benchmark
cpu: AMD EPYC 7763 64-Core Processor                
Benchmark2_AsaiYusuke_JSONPath_reuseBuffer-4   	 1000000	      1130 ns/op	      80 B/op	       2 allocs/op
Benchmark2_AsaiYusuke_JSONPath-4               	 1000000	      1169 ns/op	      96 B/op	       3 allocs/op
Benchmark2_ohler55_OjG_jp-4                    	  312778	      3794 ns/op	    6200 B/op	      37 allocs/op
Benchmark2_bhmj_JSON_Slice-4                   	   75607	     15887 ns/op	    1784 B/op	      38 allocs/op
Benchmark2_Spyzhov_Abstract_JSON-4             	   77376	     16420 ns/op	    5480 B/op	     223 allocs/op
PASS
ok  	github.com/AsaiYusuke/jsonpath_benchmark	5.965s

```

## Conclusion

This benchmark highlights significant performance differences among popular JSONPath libraries.
Developers can use these results as a reference when selecting a library, particularly when performance is critical.
For real-world use, consider running benchmarks tailored to specific datasets and queries.
