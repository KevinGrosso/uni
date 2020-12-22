package main

import (
	"testing"
)

func sliceEqual(x []int, y []int) bool{
	if len(x) != len(y) {
		return false
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func TestSort(t *testing.T) {
	type testCase struct {
		input []int
		expected []int
	}

	cases := []testCase {
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{-1, -2, 10, 11, 71}, []int{-2, -1, 10, 11, 71}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{20, 10, 20, 10, -1, 100}, []int{-1, 10, 10, 20, 20, 100}},
	}

	for _, c := range cases {
		output := make([]int, len(c.input))
		copy(output, c.input)
		sort(output)
		if !sliceEqual(output, c.expected) {
			t.Errorf("Input: %v - Output: %v - Expected: %v\n", c.input, output, c.expected)
		}
	}
}
