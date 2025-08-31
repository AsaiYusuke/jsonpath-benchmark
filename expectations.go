package jsonpath_benchmark

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type BenchExpect struct {
	Values []any
}

func (e *BenchExpect) validateSlice(gotValues []any) (bool, string) {
	if e == nil || e.Values == nil {
		return true, ""
	}
	wantValues := e.Values

	if reflect.DeepEqual(gotValues, wantValues) {
		return true, ""
	}

	return false, fmt.Sprintf("mismatch: got=%s want=%s", jsonStringify(gotValues), jsonStringify(wantValues))
}

func jsonStringify(value any) string {
	b, err := json.Marshal(value)
	if err != nil {
		return fmt.Sprintf("%#v", value)
	}
	return string(b)
}
