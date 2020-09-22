package handler_test

import "testing"

// SkipThis is a function.
func SkipThis(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test")
	}
}
