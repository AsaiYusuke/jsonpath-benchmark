package jsonpath_benchmark

import (
	"testing"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
)

func Execute_ohler55_OjG_jp(b *testing.B, srcJSON string, jsonPath string) {
	b.Helper()

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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Get(obj)
	}
}
