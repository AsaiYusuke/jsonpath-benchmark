# JSONPath Benchmark

[![Benchmark JSONPath](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml/badge.svg)](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml)

This project benchmarks multiple JSONPath libraries written in Go.
It evaluates performance in specific scenarios and offers guidance for selecting an appropriate library.

## Contents

- [Overview](#overview)
- [Libraries Benchmarked](#libraries-benchmarked)
- [Simple Query Performance](#simple-query-performance)
- [Complex Query Performance](#complex-query-performance)
- [Support Matrix](#support-matrix)
- [Conclusion](#conclusion)
- [Reproduce Locally](#reproduce-locally)
- [License](#license)

## Overview

Results are generated via GitHub Actions for consistency and automation.

The benchmark focuses on the per-operation cost after initial setup is complete.
Preparatory steps, such as parsing or preloading data, are excluded from the measurements.
Instead, it emphasizes scenarios with intensive looping or query execution within the main operation.

Results may vary with factors such as input structure, query complexity, and runtime environment.
Treat these benchmarks as a general reference and re-evaluate them periodically.

Notes on metrics:

- Time: ns/op (lower is better)
- Memory: B/op (lower is better)
- Allocations: allocs/op (lower is better)

## Libraries Benchmarked

The following libraries are included in this benchmark:

- [AsaiYusuke/JSONPath](https://github.com/AsaiYusuke/jsonpath)
- [PaesslerAG/JSONPath](https://github.com/PaesslerAG/jsonpath)
- [bhmj/JSONSlice](https://github.com/bhmj/jsonslice)
- [ohler55/OjG](https://github.com/ohler55/ojg)
- [oliveagle/JsonPath](https://github.com/oliveagle/jsonpath)
- [Spyzhov/Abstract JSON](https://github.com/spyzhov/ajson)
- [vmware-labs/YAML JSONPath](https://github.com/vmware-labs/yaml-jsonpath)

## Simple Query Performance

JSONPath:

``` text
$.store.book[0].price
```

Summary:

- All listed libraries support this query, so results are directly comparable.
- With buffer reuse, `AsaiYusuke/JSONPath` is the fastest; with per-op allocation, it ranks second.

Performance Detail:

|  Rank  | Library                     |   Time (ns/op) |   Memory (B/op) |   Allocations (allocs/op) |   Relative speed (fastest = 1x) |
|:------:|:----------------------------|---------------:|----------------:|--------------------------:|--------------------------------:|
|   1    | AsaiYusuke/JSONPath (reuse) |          62.40 |               0 |                         0 |                           1.00x |
|   2    | oliveagle/JsonPath          |          69.32 |               0 |                         0 |                           1.11x |
|   3    | AsaiYusuke/JSONPath         |         105.70 |              16 |                         1 |                           1.69x |
|   4    | ohler55/OjG (jp)            |         358.70 |            1168 |                         2 |                           5.75x |
|   5    | PaesslerAG/JSONPath         |         397.90 |             208 |                         7 |                           6.38x |
|   6    | vmware-labs/YAML JSONPath   |         908.60 |             464 |                        28 |                          14.56x |
|   7    | bhmj/JSONSlice              |        1279.00 |              24 |                         1 |                          20.50x |
|   8    | Spyzhov/ajson               |        1385.00 |             472 |                        25 |                          22.20x |

![Simple query benchmark (ns/op)](assets/bench_chart_simple.svg)

## Complex Query Performance

JSONPath:

``` text
$..book[?(@.price > $.store.bicycle.price)]
```

Summary:

- This query exercises recursive descent and filters; only a subset of libraries support it.
- Among those, `AsaiYusuke/JSONPath` delivered the best performance.

Performance Detail:

|  Rank  | Library                     |   Time (ns/op) |   Memory (B/op) |   Allocations (allocs/op) |   Relative speed (fastest = 1x) |
|:------:|:----------------------------|---------------:|----------------:|--------------------------:|--------------------------------:|
|   1    | AsaiYusuke/JSONPath (reuse) |        1097.00 |              80 |                         2 |                           1.00x |
|   2    | AsaiYusuke/JSONPath         |        1162.00 |              96 |                         3 |                           1.06x |
|   3    | ohler55/OjG (jp)            |        3889.00 |            6200 |                        37 |                           3.55x |
|   4    | Spyzhov/ajson               |       15687.00 |            5480 |                       223 |                          14.30x |
|   5    | bhmj/JSONSlice              |       15894.00 |            1784 |                        38 |                          14.49x |

![Complex query benchmark (ns/op)](assets/bench_chart_complex.svg)

## Support Matrix

| Library                     | Simple query   | Complex query   |
|:----------------------------|:---------------|:----------------|
| AsaiYusuke/JSONPath (reuse) | ✅              | ✅               |
| AsaiYusuke/JSONPath         | ✅              | ✅               |
| PaesslerAG/JSONPath         | ✅              | ❌               |
| bhmj/JSONSlice              | ✅              | ✅               |
| ohler55/OjG (jp)            | ✅              | ✅               |
| oliveagle/JsonPath          | ✅              | ❌               |
| Spyzhov/ajson               | ✅              | ✅               |
| vmware-labs/YAML JSONPath   | ✅              | ❌               |

## Conclusion

This benchmark compared several popular JSONPath libraries in Go and highlighted notable performance differences.
Interestingly, the simple query showed a wider performance spread than the complex one, suggesting that implementation details and variations in query syntax handling can directly impact execution speed.
Therefore, a practical evaluation should consider both feature support and raw performance.
For selecting a library in production, we strongly recommend running benchmarks tailored to your own datasets and query patterns.

## Reproduce Locally

Benchmarks are executed in GitHub Actions for consistency.
For the exact steps and current outputs, check the Actions tab and the workflow logs.
If you prefer to run locally, follow the same sequence defined in the workflow file (see [.github/workflows/build.yml](.github/workflows/build.yml)).

## License

This project is distributed under the terms of the MIT License. See [LICENSE](LICENSE) for details.
