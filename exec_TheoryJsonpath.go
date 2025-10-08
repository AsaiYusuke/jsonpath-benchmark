package jsonpath_benchmark

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/theory/jsonpath"
)

func Execute_theory_jsonpath(b *testing.B, srcJSON string, jsonPath string, expect *BenchExpect) {

	var src any
	if err := json.Unmarshal([]byte(srcJSON), &src); err != nil {
		b.Skip(err)
		return
	}

	path, err := jsonpath.Parse(jsonPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	nodes := path.Select(src)
	var result []any
	for _, node := range nodes {
		result = append(result, node)
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
		path.Select(src)
	}
}
