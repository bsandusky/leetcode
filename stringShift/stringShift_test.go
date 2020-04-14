package stringShift

import (
	"reflect"
	"testing"
)

var tts = []struct {
	str   string
	shift [][]int
	out   string
}{
	{
		str:   "abc",
		shift: [][]int{[]int{0, 1}, []int{1, 2}},
		out:   "cab",
	},
	{
		str:   "abcdefg",
		shift: [][]int{[]int{1, 1}, []int{1, 1}, []int{0, 2}, []int{1, 3}},
		out:   "efgabcd",
	},
	{
		str:   "abc",
		shift: [][]int{[]int{1, 1}, []int{0, 2}},
		out:   "bca",
	},
	{
		str:   "yisxjwry",
		shift: [][]int{[]int{1, 8}, []int{1, 4}, []int{1, 3}, []int{1, 6}, []int{0, 6}, []int{1, 4}, []int{0, 2}, []int{0, 1}},
		out:   "yisxjwry",
	},
}

func TestStringShift(t *testing.T) {
	for _, tt := range tts {
		got := stringShift(tt.str, tt.shift)
		if !reflect.DeepEqual(got, tt.out) {
			t.Errorf("Input string: %s, Input shift: %v | Expected: %v, Got: %v", tt.str, tt.shift, tt.out, got)
		}
	}
}
