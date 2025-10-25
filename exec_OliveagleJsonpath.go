package jsonpath_benchmark

import (
	"encoding/json"
	"testing"

	"github.com/oliveagle/jsonpath"
)

func Execute_oliveagle_JsonPath(b *testing.B, srcJSON string, jsonPath string, expect *BenchExpect) {

	var src any
	if err := json.Unmarshal([]byte(srcJSON), &src); err != nil {
		b.Skip(err)
		return
	}

	pat, err := jsonpath.Compile(jsonPath)
	if err != nil {
		b.Skip(err)
		return
	}
	value, err := pat.Lookup(src)
	if err != nil {
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
		result = []any{result}
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
		pat.Lookup(src)
	}
}
