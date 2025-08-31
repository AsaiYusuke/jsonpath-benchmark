package jsonpath_benchmark

import (
	"encoding/json"
	"testing"

	"github.com/AsaiYusuke/jsonpath/v2"
)

func Execute_AsaiYusuke_JSONPath(b *testing.B, srcJSON string, jsonPath string, expect *BenchExpect) {

	var src any
	if err := json.Unmarshal([]byte(srcJSON), &src); err != nil {
		b.Skip(err)
		return
	}

	parserFunc, err := jsonpath.Parse(jsonPath)
	if err != nil {
		b.Skip(err)
		return
	}
	result, err := parserFunc(src)
	if err != nil {
		b.Skip(err)
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
