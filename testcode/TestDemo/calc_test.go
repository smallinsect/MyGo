package testdemo

import "testing"

// TestAdd is
func TestAdd(t *testing.T) {
	r := Add(11, 22)
	if r != 33 {
		t.Fatalf("%d, %d", 33, r)
	}
	t.Logf("test Add succ")
}

// TestSub is
func TestSub(t *testing.T) {
	r := Sub(11, 22)
	if r != -11 {
		t.Fatalf("%d, %d", -11, r)
	}
	t.Logf("test Sub succ")
}
