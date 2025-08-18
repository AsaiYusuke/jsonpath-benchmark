package jsonpath_benchmark

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/PaesslerAG/jsonpath"
)

func Execute_PaesslerAG_JSONPath(b *testing.B, srcJSON string, jsonPath string) {
	b.Helper()

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

	for b.Loop() {
		if _, err := eval(context.Background(), src); err != nil {
			b.Skip(err)
			return
		}
	}
}
