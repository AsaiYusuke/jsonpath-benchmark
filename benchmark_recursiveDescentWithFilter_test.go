package jsonpath_benchmark

import (
	"testing"
)

var jsonPath_recursiveDescentWithFilter string = `$..book[?(@.price > $.store.bicycle.price)]`
var expect_recursiveDescentWithFilter = &BenchExpect{Values: []any{map[string]any{
	"category": "fiction",
	"author":   "J. R. R. Tolkien",
	"title":    "The Lord of the Rings",
	"isbn":     "0-395-19395-8",
	"price":    22.99,
}}}

func Benchmark2_AsaiYusuke_JSONPath_reuseBuffer(b *testing.B) {
	Execute_AsaiYusuke_JSONPath_reuseBuffer(b, srcJSON, jsonPath_recursiveDescentWithFilter, expect_recursiveDescentWithFilter)
}

func Benchmark2_AsaiYusuke_JSONPath(b *testing.B) {
	Execute_AsaiYusuke_JSONPath(b, srcJSON, jsonPath_recursiveDescentWithFilter, expect_recursiveDescentWithFilter)
}

func Benchmark2_ohler55_OjG_jp(b *testing.B) {
	Execute_ohler55_OjG_jp(b, srcJSON, jsonPath_recursiveDescentWithFilter, expect_recursiveDescentWithFilter)
}

func Benchmark2_bhmj_JSON_Slice(b *testing.B) {
	Execute_bhmj_JSON_Slice(b, srcJSON, jsonPath_recursiveDescentWithFilter, expect_recursiveDescentWithFilter)
}

func Benchmark2_Spyzhov_Abstract_JSON(b *testing.B) {
	Execute_Spyzhov_Abstract_JSON(b, srcJSON, jsonPath_recursiveDescentWithFilter, expect_recursiveDescentWithFilter)
}

func Benchmark2_vmware_labs_YAML_JSONPath(b *testing.B) {
	Execute_vmware_labs_YAML_JSONPath(b, srcJSON, jsonPath_recursiveDescentWithFilter, expect_recursiveDescentWithFilter)
}

func Benchmark2_oliveagle_JsonPath(b *testing.B) {
	Execute_oliveagle_JsonPath(b, srcJSON, jsonPath_recursiveDescentWithFilter, expect_recursiveDescentWithFilter)
}

func Benchmark2_PaesslerAG_JSONPath(b *testing.B) {
	Execute_PaesslerAG_JSONPath(b, srcJSON, jsonPath_recursiveDescentWithFilter, expect_recursiveDescentWithFilter)
}
