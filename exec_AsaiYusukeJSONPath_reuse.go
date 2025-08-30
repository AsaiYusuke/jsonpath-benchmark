package jsonpath_benchmark

import (
	"encoding/json"
	"testing"

	"github.com/AsaiYusuke/jsonpath/v2"
)

func Execute_AsaiYusuke_JSONPath_reuseBuffer(b *testing.B, srcJSON string, jsonPath string) {
	b.Helper()

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

	buf := make([]any, 0, 256)
	args := []*[]any{&buf}

	for b.Loop() {
		if _, err := parserFunc(src, args...); err != nil {
			b.Skip(err)
			return
		}
	}
}
