package productExceptSelf

import (
	"reflect"
	"testing"
)

var tts = []struct {
	in  []int
	out []int
}{
	{in: []int{1, 2, 3, 4}, out: []int{24, 12, 8, 6}},
}

func TestProductExceptSelf(t *testing.T) {
	for _, tt := range tts {
		got := productExceptSelf(tt.in)
		if !reflect.DeepEqual(tt.out, got) {
			t.Errorf("Input: %v | Expected: %v, Got: %v", tt.in, tt.out, got)
		}
	}
}
