package iwalkdeep_test

import (
	. "github.com/oderwat/iwalkdeep"
	"testing"
)

// Testing IWalkDeep can not rely on any order
// of the returned data but of that it is "deep sorted"
// the directories in a single folder can be any order
// We also ignore any regular files!

func TestDeepFirst(t *testing.T) {

	// Testing the Deep First
	results := []string{
		"testdata/A/a/b/c",
		"testdata/A/a/b",
		"testdata/A/a",
		"testdata/A"}

	idx := 0
	for dir := range IWalkDeep("testdata/A") {
		if dir != results[idx] {
			t.Errorf("%d: %s != %s", idx, dir, results[idx])
		}
		idx++
	}

}

func TestDirB(t *testing.T) {

	// Testing for multiple dirs and deep first
	results := map[string]bool{
		"testdata/B/a": true,
		"testdata/B/b": true,
		"testdata/B/c": true}

	idx := 0
	for dir := range IWalkDeep("testdata/B") {
		if idx > 3 {
			t.Error("More entries found as expected")
		} else if idx == 3 {
			if dir != "testdata/B" {
				t.Error("Final directory was not as expected")
			}
			if len(results) != 0 {
				t.Error("Some Dir entries where missing.")
			}
		} else if results[dir] {
			delete(results, dir)
		} else {
			t.Errorf("%d: %s not in expected results", idx, dir)
		}
		idx++
	}
}

func TestDirAll(t *testing.T) {

	// Test if we really walk all dirs
	results := map[string]bool{
		"testdata/A/a/b/c": true,
		"testdata/A/a/b":   true,
		"testdata/A/a":     true,
		"testdata/A":       true,
		"testdata/B/a":     true,
		"testdata/B/b":     true,
		"testdata/B/c":     true,
		"testdata/B":       true,
		"testdata/C":       true,
		"testdata":         true}

	idx := 0
	for dir := range IWalkDeep("testdata") {
		if results[dir] {
			delete(results, dir)
		} else {
			t.Errorf("%d: %s not in expected results", idx, dir)
		}
		idx++
	}
	if len(results) != 0 {
		t.Error("Some Dir entries where missing.")
	}
}
