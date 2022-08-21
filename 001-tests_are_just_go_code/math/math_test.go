package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	want := 11
	got := Sum([]int{10, -2, 3})

	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}

func TestAdd(t *testing.T) {
	want := 5
	got := Add(2, 3)

	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}
