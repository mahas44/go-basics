package test

import "testing"

func TestMySum(t *testing.T) {
	x := MySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}

func TestMyAverage(t *testing.T) {
	x := MyAverage([]float64{4, 5, 6, 7})
	if x != 5.5 {
		t.Error("Expected", 5.5, "Got", x)
	}
}
