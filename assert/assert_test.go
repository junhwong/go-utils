package assert

import "testing"

func TestAssert(t *testing.T) {
	// var err error = New("a", "f")
	// t.Logf("%-v", err)
	DEBUG = false
	Assert(false, "f")
}
