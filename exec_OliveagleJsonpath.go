package jsonpath_benchmark

import (
	"encoding/json"
	"testing"

	"github.com/oliveagle/jsonpath"
)

func Execute_oliveagle_JsonPath(b *testing.B, srcJSON string, jsonPath string) {
	b.Helper()

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

	for b.Loop() {
		if _, err := pat.Lookup(src); err != nil {
			b.Skip(err)
			return
		}
	}
}
