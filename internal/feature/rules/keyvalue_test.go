package rules

import (
	"fmt"
	"testing"
)

func TestQueryString(t *testing.T) {

	values := map[string]string{
		"param1": "value1",
		"param2": "value2",
		"param3": "value3",
		"param4": "value4",
		"param5": "value5",
		"param6": "value6",
	}
	var tests = []struct {
		qparam   KeyValue
		values   map[string]string
		expected bool
	}{
		{KeyValue{Key: "param1", Value: []string{"value0"}}, values, false},
		{KeyValue{Key: "param1", Value: []string{"value1"}}, values, true},
		{KeyValue{Key: "param10", Value: []string{"xyz"}}, values, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("version: %s", tt.qparam.Key)
		t.Run(testname, func(t *testing.T) {
			var rq RuleKeyValue
			ok := rq.evaluate(tt.qparam, tt.values)
			if ok != tt.expected {
				t.Errorf("got %t, want %t", ok, tt.expected)
			}
		})
	}

}
