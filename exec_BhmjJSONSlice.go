package jsonpath_benchmark

import (
	"testing"

	"github.com/bhmj/jsonslice"
)

func Execute_bhmj_JSON_Slice(b *testing.B, srcJSON string, jsonPath string) {
	b.Helper()

	var src = []byte(srcJSON)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := jsonslice.Get(src, jsonPath); err != nil {
			b.Skip(err)
			return
		}
	}
}
