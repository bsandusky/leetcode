package moveZeroes

import (
	"reflect"
	"testing"
)

var tts = []struct {
	in  []int
	out []int
}{
	{in: []int{0, 1, 0, 3, 12}, out: []int{1, 3, 12, 0, 0}},
	{in: []int{0, 0, 1}, out: []int{1, 0, 0}},
}

func TestMoveZeroes(t *testing.T) {
	for _, tt := range tts {
		got := moveZeroes(tt.in)
		if !reflect.DeepEqual(tt.out, got) {
			t.Errorf("Input: %v | Expected: %v, Got: %v", tt.in, tt.out, got)
		}
	}
}
