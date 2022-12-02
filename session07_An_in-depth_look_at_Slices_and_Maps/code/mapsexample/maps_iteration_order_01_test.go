package mapsexample

import (
	"sort"
	"testing"
)

func TestRandomIterationOrder(t *testing.T) {
	var m = map[int]string{
		100: "hundred",
		1:   "one",
		10:  "ten",
	}
	for k, v := range m {
		t.Logf("Key = %d; Value = %q\n", k, v)
	}
}

func TestIterationOrderPreserved(t *testing.T) {
	var m = map[int]string{
		100: "hundred",
		1:   "one",
		10:  "ten",
	}
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		t.Logf("Key = %d; Value = %q\n", k, m[k])
	}
}
