# JSONPath Benchmark

[![Benchmark JSONPath](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml/badge.svg)](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml)

This project benchmarks multiple JSONPath libraries written in Go.
It evaluates performance in specific scenarios and offers guidance for selecting an appropriate library.

## Contents

- [Overview](#overview)
- [Libraries Benchmarked](#libraries-benchmarked)
- [Results: Simple Query](#results-simple-query)
- [Results: Complex Query](#results-complex-query)
- [Capability Overview](#capability-overview)
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

- All listed libraries support this query, so results are directly comparable.
- With buffer reuse, AsaiYusuke/JSONPath is the fastest; with per-op allocation, it ranks second.

### Simple query results

{% include 'assets/bench_table_simple.md' %}

![Simple query benchmark (ns/op)](assets/bench_chart_simple.svg)

## Results: Complex Query

JSONPath:

``` text
$..book[?(@.price > $.store.bicycle.price)]
```

Performance Summary:

- This query exercises recursive descent and filters; only a subset of libraries support it.
- Among those, AsaiYusuke/JSONPath delivered the best performance.

### Complex query results

{% include 'assets/bench_table_complex.md' %}

![Complex query benchmark (ns/op)](assets/bench_chart_complex.svg)

## Capability Overview

{% include 'assets/compat_matrix.md' %}

## Conclusion

This benchmark compared several popular JSONPath libraries in Go and highlighted notable performance differences.
Interestingly, the simple query showed a wider performance spread than the complex one, suggesting that implementation details and variations in query syntax handling can directly impact execution speed.
Therefore, a practical evaluation should consider both feature support and raw performance.
For selecting a library in production, we strongly recommend running benchmarks tailored to your own datasets and query patterns.

## Reproduce Locally

Benchmarks are executed in GitHub Actions for consistency. For the exact steps and current outputs, check the Actions tab and the workflow logs. If you prefer to run locally, follow the same sequence defined in the workflow file (see `.github/workflows/build.yml`).

## License

This project is distributed under the terms of the MIT License. See [LICENSE](LICENSE) for details.
