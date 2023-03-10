package jsonpath_benchmark

import (
	"testing"

	"github.com/spyzhov/ajson"
)

func Execute_Spyzhov_Abstract_JSON(b *testing.B, srcJSON string, jsonPath string) {
	b.Helper()

	json := []byte(srcJSON)

	root, _ := ajson.Unmarshal(json)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := root.JSONPath(jsonPath); err != nil {
			b.Skip(err)
			return
		}
	}
}
