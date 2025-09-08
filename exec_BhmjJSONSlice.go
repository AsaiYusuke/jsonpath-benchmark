package jsonpath_benchmark

import (
	"encoding/json"
	"testing"

	"github.com/bhmj/jsonslice"
)

func Execute_bhmj_JSON_Slice(b *testing.B, srcJSON string, jsonPath string, expect *BenchExpect) {

	var src = []byte(srcJSON)
	binaryResult, err := jsonslice.Get(src, jsonPath)
	if err != nil {
		b.Skip(err)
		return
	}

	var value any
	if err := json.Unmarshal(binaryResult, &value); err != nil {
		b.Skip(err)
		return
	}

	result, ok := value.([]any)
	if ok {
		if len(result) == 1 {
			_result, ok := result[0].([]any)
			if ok {
				result = _result
			}
		}
	} else {
		result = []any{value}
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
		if _, err := jsonslice.Get(src, jsonPath); err != nil {
			b.Skip(err)
			return
		}
	}
}
