package jsonpath_benchmark

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/PaesslerAG/jsonpath"
)

func Execute_PaesslerAG_JSONPath(b *testing.B, srcJSON string, jsonPath string, expect *BenchExpect) {

	var src any
	if err := json.Unmarshal([]byte(srcJSON), &src); err != nil {
		b.Skip(err)
		return
	}

	eval, err := jsonpath.New(jsonPath)
	if err != nil {
		b.Skip(err)
		return
	}
	value, err := eval(context.Background(), src)
	if err != nil {
		b.Skip(err)
		return
	}
	result := []any{value}
	if ok, reason := expect.validateSlice(result); !ok {
		if reason != "" {
			b.Skipf("precheck failed: %s", reason)
		} else {
			b.Skip("precheck failed")
		}
		return
	}

	for b.Loop() {
		if _, err := eval(context.Background(), src); err != nil {
			b.Skip(err)
			return
		}
	}
}
