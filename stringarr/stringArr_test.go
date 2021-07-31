package stringarr

import (
	"strings"
	"testing"
)

func TestStringArr(t *testing.T) {

	var strs = []string{"watermelon", "peach", "apple", "pear", "orange", "plum"}

	t.Run("Index", func(t *testing.T) {
		if Index(strs, "pear") != 3 {
			t.Errorf("Index failed")
		}
	})
	t.Run("Include", func(t *testing.T) {
		if Include(strs, "melon") != false {
			t.Errorf("Include failed")
		}
	})
	t.Run("Any", func(t *testing.T) {
		res := Any(strs, func(v string) bool {
			return strings.HasPrefix(v, "p")
		})
		if res != true {
			t.Errorf("Any failed")
		}
	})
	t.Run("All", func(t *testing.T) {
		res := All(strs, func(v string) bool {
			return strings.HasPrefix(v, "p")
		})
		if res != false {
			t.Errorf("All failed")
		}
	})
	t.Run("Filter", func(t *testing.T) {
		res := Filter(strs, func(v string) bool {
			return strings.Contains(v, "e")
		})
		if len(res) == 4 {
			t.Errorf("Filter failed")
		}
	})
	t.Run("Map", func(t *testing.T) {
		res := Map(strs, strings.ToUpper)
		if len(res) != 6 {
			t.Errorf("Map failed")
		}
	})

}
