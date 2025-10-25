package jsonpath_benchmark

import (
	"encoding/json"
	"testing"

	"github.com/vmware-labs/yaml-jsonpath/pkg/yamlpath"
	"gopkg.in/yaml.v3"
)

func Execute_vmware_labs_YAML_JSONPath(b *testing.B, srcJSON string, jsonPath string, expect *BenchExpect) {

	var src any
	if err := json.Unmarshal([]byte(srcJSON), &src); err != nil {
		b.Skip(err)
		return
	}

	yamlData, err := yaml.Marshal(src)
	if err != nil {
		b.Skip(err)
		return
	}

	var srcYamlNode yaml.Node
	if err := yaml.Unmarshal(yamlData, &srcYamlNode); err != nil {
		b.Skip(err)
		return
	}

	path, err := yamlpath.NewPath(jsonPath)
	if err != nil {
		b.Skip(err)
		return
	}
	nodes, err := path.Find(&srcYamlNode)
	if err != nil {
		b.Skip(err)
		return
	}
	var result []any
	for _, node := range nodes {
		var v any
		if err := node.Decode(&v); err == nil {
			result = append(result, v)
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
		path.Find(&srcYamlNode)
	}
}
