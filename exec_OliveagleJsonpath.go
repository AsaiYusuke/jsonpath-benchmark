package jsonpath_benchmark

import (
	"encoding/json"
	"testing"

	"github.com/oliveagle/jsonpath"
)

func Execute_oliveagle_JsonPath(b *testing.B, srcJSON string, jsonPath string) {
	b.Helper()

	var src interface{}
	if err := json.Unmarshal([]byte(srcJSON), &src); err != nil {
		b.Skip(err)
		return
	}

	pat, err := jsonpath.Compile(jsonPath)
	if err != nil {
		b.Skip(err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := pat.Lookup(src); err != nil {
			b.Skip(err)
			return
		}
	}
}
