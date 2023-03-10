package jsonpath_benchmark

import (
	"testing"
)

var jsonPath_threeLevelsWithIndex string = `$.store.book[0].price`

func Benchmark1_oliveagle_JsonPath(b *testing.B) {
	Execute_oliveagle_JsonPath(b, srcJSON, jsonPath_threeLevelsWithIndex)
}

func Benchmark1_AsaiYusuke_JSONPath(b *testing.B) {
	Execute_AsaiYusuke_JSONPath(b, srcJSON, jsonPath_threeLevelsWithIndex)
}

func Benchmark1_ohler55_OjG_jp(b *testing.B) {
	Execute_ohler55_OjG_jp(b, srcJSON, jsonPath_threeLevelsWithIndex)
}

func Benchmark1_PaesslerAG_JSONPath(b *testing.B) {
	Execute_PaesslerAG_JSONPath(b, srcJSON, jsonPath_threeLevelsWithIndex)
}

func Benchmark1_vmware_labs_YAML_JSONPath(b *testing.B) {
	Execute_vmware_labs_YAML_JSONPath(b, srcJSON, jsonPath_threeLevelsWithIndex)
}

func Benchmark1_bhmj_JSON_Slice(b *testing.B) {
	Execute_bhmj_JSON_Slice(b, srcJSON, jsonPath_threeLevelsWithIndex)
}

func Benchmark1_Spyzhov_Abstract_JSON(b *testing.B) {
	Execute_Spyzhov_Abstract_JSON(b, srcJSON, jsonPath_threeLevelsWithIndex)
}
