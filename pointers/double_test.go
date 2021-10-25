package pointers_test

import (
	"testing"

	"github.com/Drew-Kimberly/go-exploration/pointers"
)

func TestDouble(t *testing.T) {
	t.Parallel()
	x := 12
	want := 24
	pointers.Double(&x)
	if want != x {
		t.Errorf("want %d, got 	%d", want, x)
	}
}

func TestDoubleNilInput(t *testing.T) {
	t.Parallel()
	pointers.Double(nil)
}
