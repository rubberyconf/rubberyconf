package feature

import (
	"fmt"
	"testing"
)

func TestCheckVersions(t *testing.T) {
	var tests = []struct {
		obj      []string
		sent     string
		expected bool
	}{
		{[]string{"1.0.1", "2.0.1"}, "1.0.1", true},
		{[]string{"1.0.1"}, "1.0.1", true},
		{[]string{"1.0.1", "2.1.1"}, "2.2.1", false},
		{[]string{">1.0.1"}, "1.0.1", false},
		{[]string{">1.0.1"}, "1.0.2", true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("version: %s", tt.sent)
		t.Run(testname, func(t *testing.T) {
			ok := versionCheck(tt.obj, tt.sent)
			if ok != tt.expected {
				t.Errorf("got %t, want %t", ok, tt.expected)
			}
		})
	}

}
