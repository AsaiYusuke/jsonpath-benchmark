# JSONPath Benchmark

[![Benchmark JSONPath](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml/badge.svg)](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml)

This project provides a comprehensive benchmark comparing multiple JSONPath libraries implemented in the Go programming language.
It aims to evaluate performance under specific scenarios and offer insights for developers choosing an appropriate library for their use case.

## Overview

Results are generated using GitHub Actions to ensure consistency and automation.

The focus of the benchmark is on the cost per operation after the initial setup is complete.
This means that preparatory steps, such as parsing or preloading data, are excluded from the measurements.
Instead, the benchmark emphasizes scenarios involving intensive looping or query execution within the main operation.

The results may vary depending on several factors, such as the structure of the input data, the complexity of the JSONPath queries, and the specific runtime environment.
For this reason, these benchmarks should be considered as a general reference and re-evaluated periodically.

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

This query was supported by all libraries included in the benchmark, enabling a direct performance comparison across all of them.
The fastest performance was achieved by `oliveagle/JsonPath`, with my library ranking second.

``` bash
goos: linux
goarch: amd64
pkg: github.com/AsaiYusuke/jsonpath_benchmark
cpu: AMD EPYC 7763 64-Core Processor                
Benchmark1_AsaiYusuke_JSONPath_reuseBuffer-4   	19454383	        61.15 ns/op	       0 B/op	       0 allocs/op
Benchmark1_oliveagle_JsonPath-4                	16255122	        73.63 ns/op	       0 B/op	       0 allocs/op
Benchmark1_AsaiYusuke_JSONPath-4               	12125500	        97.98 ns/op	      16 B/op	       1 allocs/op
Benchmark1_ohler55_OjG_jp-4                    	 3435986	       351.8 ns/op	    1168 B/op	       2 allocs/op
Benchmark1_PaesslerAG_JSONPath-4               	 2968284	       403.9 ns/op	     208 B/op	       7 allocs/op
Benchmark1_vmware_labs_YAML_JSONPath-4         	 1325276	       906.2 ns/op	     464 B/op	      28 allocs/op
Benchmark1_bhmj_JSON_Slice-4                   	  931850	      1270 ns/op	      24 B/op	       1 allocs/op
Benchmark1_Spyzhov_Abstract_JSON-4             	  854456	      1426 ns/op	     472 B/op	      25 allocs/op
PASS
ok  	github.com/AsaiYusuke/jsonpath_benchmark	9.593s

```

## Results: Complex Query

JSONPath:

``` text
$..book[?(@.price > $.store.bicycle.price)]
```

Performance Summary:

This query involves more complex syntax, and only a subset of the libraries were able to process it.
Among these, my library demonstrated the best performance.

``` bash
goos: linux
goarch: amd64
pkg: github.com/AsaiYusuke/jsonpath_benchmark
cpu: AMD EPYC 7763 64-Core Processor                
Benchmark2_AsaiYusuke_JSONPath_reuseBuffer-4   	 1000000	      1123 ns/op	      80 B/op	       2 allocs/op
Benchmark2_AsaiYusuke_JSONPath-4               	 1000000	      1170 ns/op	      96 B/op	       3 allocs/op
Benchmark2_ohler55_OjG_jp-4                    	  323191	      3670 ns/op	    6200 B/op	      37 allocs/op
Benchmark2_vmware_labs_YAML_JSONPath-4         	  278862	      4276 ns/op	    4416 B/op	     136 allocs/op
Benchmark2_bhmj_JSON_Slice-4                   	   75988	     15590 ns/op	    1784 B/op	      38 allocs/op
Benchmark2_Spyzhov_Abstract_JSON-4             	   77359	     15470 ns/op	    5480 B/op	     223 allocs/op
PASS
ok  	github.com/AsaiYusuke/jsonpath_benchmark	7.060s

```

## Conclusion

This benchmark highlights significant performance differences among popular JSONPath libraries.
Developers can use these results as a reference when selecting a library, particularly when performance is critical.
For real-world use, consider running benchmarks tailored to your specific datasets and queries.
