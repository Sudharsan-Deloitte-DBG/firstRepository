package main

import "testing"

func TestAdd(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{30, 21}, 51},
		test{[]int{41, -10}, 31},
		test{[]int{-11, -43}, -54},
	}

	sum := 0
	for _, v := range tests {
		sum = add(v.data...)
		if sum != v.answer {
			t.Errorf("Addition failed\nExpected: %v\nGot: %v", v.answer, sum)
		}
	}
}
