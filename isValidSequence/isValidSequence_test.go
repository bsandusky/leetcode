package isValidSequence

import "testing"

var tts = []struct {
	in  *TreeNode
	arr []int
	out bool
}{
	{
		in: &TreeNode{Val: 0,
			Left: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 0, Right: &TreeNode{Val: 1}},
				Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}}},
			Right: &TreeNode{Val: 0,
				Right: &TreeNode{Val: 0, Left: &TreeNode{Val: 0}}},
		},
		arr: []int{0, 1, 0, 1},
		out: true,
	},
}

func TestIsValidSequence(t *testing.T) {

	for _, tt := range tts {
		got := isValidSequence(tt.in, tt.arr)

		if tt.out != got {
			t.Errorf("Expected: %v; Got: %v", tt.out, got)
		}
	}
}
