package jsonpath_benchmark

import (
	"testing"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
)

func Execute_ohler55_OjG_jp(b *testing.B, srcJSON string, jsonPath string, expect *BenchExpect) {

	obj, err := oj.ParseString(srcJSON)
	if err != nil {
		b.Skip(err)
		return
	}

	x, err := jp.ParseString(jsonPath)
	if err != nil {
		b.Skip(err)
		return
	}
	result := x.Get(obj)
	if ok, reason := expect.validateSlice(result); !ok {
		if reason != "" {
			b.Skipf("precheck failed: %s", reason)
		} else {
			b.Skip("precheck failed")
		}
		return
	}

	for b.Loop() {
		x.Get(obj)
	}
}
