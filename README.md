# JSONPath Benchmark

[![Benchmark JSONPath](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml/badge.svg)](https://github.com/AsaiYusuke/jsonpath-benchmark/actions/workflows/build.yml)

I benchmarked two JSONPaths using several libraries for the Go language.
Results can be checked via Github actions.

What is being measured is the cost per job for a job that loops a lot after all the prep work done.

There was a performance differences.
But if the number of queries is little, there will not be a big difference between any of them.

Also, the results will vary depending on the data entered.
So this benchmark is for information only and should be re-measured at every time.

- [AsaiYusuke/JSONPath](https://github.com/AsaiYusuke/jsonpath)
- [ohler55/OjG](https://github.com/ohler55/ojg)
- [vmware-labs/YAML JSONPath](https://github.com/vmware-labs/yaml-jsonpath)
- [bhmj/JSONSlice](https://github.com/bhmj/jsonslice)
- [Spyzhov/Abstract JSON](https://github.com/spyzhov/ajson)
- [oliveagle/JsonPath](https://github.com/oliveagle/jsonpath)
- [PaesslerAG/JSONPath](https://github.com/PaesslerAG/jsonpath)

## JSONPath for comparison with more libraries

```
JSONPath : $.store.book[0].price
```

This is the result of a JSONPath that all libraries were able to process.
oliveagle/JsonPath is fastest. My library is 2nd.

## A slightly complex JSONPath

```
JSONPath : $..book[?(@.price > $.store.bicycle.price)]
```

Libraries that can handle complex syntax limited to a few.
Among these libraries, my library is fastest.