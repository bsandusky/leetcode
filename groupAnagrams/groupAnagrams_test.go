package groupAnagrams

import (
	"reflect"
	"testing"
)

var tts = []struct {
	in  []string
	out [][]string
}{
	{
		in:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		out: [][]string{[]string{"ate", "eat", "tea"}, []string{"nat", "tan"}, []string{"bat"}},
	},
}

func TestGroupAnagrams(t *testing.T) {
	for _, tt := range tts {
		got := groupAnagrams(tt.in)
		if !reflect.DeepEqual(got, tt.out) {
			t.Errorf("Input: %v | Expected: %v, Got: %v", tt.in, tt.out, got)
		}
	}
}
