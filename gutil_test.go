package gutil

import "testing"

func TestRand(t *testing.T) {
	t.Logf("Str-> %s", RandString(43))
}
