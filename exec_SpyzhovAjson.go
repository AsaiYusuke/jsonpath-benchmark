package jsonpath_benchmark

import (
	"testing"

	"github.com/spyzhov/ajson"
)

func Execute_Spyzhov_Abstract_JSON(b *testing.B, srcJSON string, jsonPath string, expect *BenchExpect) {

	json := []byte(srcJSON)

	root, _ := ajson.Unmarshal(json)

	nodes, err := root.JSONPath(jsonPath)
	if err != nil {
		b.Skip(err)
		return
	}

	result := make([]any, 0, len(nodes))
	for _, n := range nodes {
		switch n.Type() {
		case ajson.Numeric:
			if value, err := n.GetNumeric(); err == nil {
				result = append(result, value)
			}
		case ajson.Object:
			if value, err := n.GetObject(); err == nil {
				m := map[string]any{}
				for key, child := range value {
					switch child.Type() {
					case ajson.String:
						if s, err := child.GetString(); err == nil {
							m[key] = s
						}
					case ajson.Numeric:
						if n, err := child.GetNumeric(); err == nil {
							m[key] = n
						}
					}
				}
				result = append(result, m)
			}
		}
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
		root.JSONPath(jsonPath)
	}
}
