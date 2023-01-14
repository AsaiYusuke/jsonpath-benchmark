package jsonpath_benchmark

import (
	"encoding/json"
	"testing"

	"github.com/AsaiYusuke/jsonpath"
)

func Execute_AsaiYusuke_JSONPath(b *testing.B, srcJSON string, jsonPath string) {
	b.Helper()

	var src interface{}
	if err := json.Unmarshal([]byte(srcJSON), &src); err != nil {
		b.Skip(err)
		return
	}

	parserFunc, err := jsonpath.Parse(jsonPath)
	if err != nil {
		b.Skip(err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := parserFunc(src); err != nil {
			b.Skip(err)
			return
		}
	}
}
