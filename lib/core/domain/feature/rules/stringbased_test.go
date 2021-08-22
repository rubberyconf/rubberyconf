package rules

import (
	"fmt"
	"testing"
)

func TestStringBased(t *testing.T) {

	potentialValues := []string{"value1", "value2", "value3", "value4"}
	var tests = []struct {
		value           string
		potentialValues []string
		expected        bool
	}{
		{"value2", potentialValues, true},
		{"value50", potentialValues, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("version: %s", tt.value)
		t.Run(testname, func(t *testing.T) {
			var rs StringArrayBased
			ok := rs.evaluate(tt.potentialValues, tt.value)
			if ok != tt.expected {
				t.Errorf("got %t, want %t", ok, tt.expected)
			}
		})
	}

}
