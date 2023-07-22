package test

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{3, 4, 5}, 12},
		{[]int{1, 1}, 2},
		{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		x := MySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func TestMyAverage(t *testing.T) {
	x := MyAverage([]float64{4, 5, 6, 7})
	if x != 5.5 {
		t.Error("Expected", 5.5, "Got", x)
	}
}
