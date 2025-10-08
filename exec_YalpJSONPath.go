package jsonpath_benchmark

import (
	"encoding/json"
	"testing"

	"github.com/yalp/jsonpath"
)

func Execute_Yalp_JSONPath(b *testing.B, srcJSON string, jsonPath string, expect *BenchExpect) {

	var src any
	if err := json.Unmarshal([]byte(srcJSON), &src); err != nil {
		b.Skip(err)
		return
	}

	parserFunc, err := jsonpath.Prepare(jsonPath)
	if err != nil {
		b.Skip(err)
		return
	}
	nodes, err := parserFunc(src)
	if err != nil {
		b.Skip(err)
		return
	}
	var result []any
	switch typedNodes := nodes.(type) {
	case []any:
		result = typedNodes
	case float64:
		result = []any{typedNodes}
	default:
		b.Skip("unexpected result type")
		return
	}

	if ok, reason := expect.validateSlice(result); !ok {
		if reason != "" {
			b.Skipf("precheck failed: %s", reason)
		} else {
			b.Skip("precheck failed")
		}
		return
	}

	for b.Loop() {
		if _, err := parserFunc(src); err != nil {
			b.Skip(err)
			return
		}
	}
}
