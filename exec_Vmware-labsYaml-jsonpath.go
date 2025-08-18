package jsonpath_benchmark

import (
	"testing"

	"github.com/vmware-labs/yaml-jsonpath/pkg/yamlpath"
	"gopkg.in/yaml.v3"
)

func Execute_vmware_labs_YAML_JSONPath(b *testing.B, srcJSON string, jsonPath string) {
	b.Helper()

	var n yaml.Node
	if err := yaml.Unmarshal([]byte(srcJSON), &n); err != nil {
		b.Skip(err)
		return
	}

	path, err := yamlpath.NewPath(jsonPath)
	if err != nil {
		b.Skip(err)
		return
	}

	for b.Loop() {
		if _, err := path.Find(&n); err != nil {
			b.Skip(err)
			return
		}
	}
}
