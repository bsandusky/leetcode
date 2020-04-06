package twoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	tests := []struct {
		in     []int
		target int
		out    []int
	}{
		{[]int{1, 2}, 3, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{5, 20, 7, 33}, 27, []int{1, 2}},
		{[]int{5, 20, 33, 44, 7}, 27, []int{1, 4}},
	}

	for _, test := range tests {
		got := twoSum(test.in, test.target)
		if !reflect.DeepEqual(got, test.out) {
			t.Errorf("Got: %v, Expected: %v\n", got, test.out)
		}
	}

}
