package diameterOfBinaryTree

import "testing"

var tts = []struct {
	in  *TreeNode
	out int
}{
	{
		in:  &TreeNode{Val: 1, Right: &TreeNode{Val: 3}, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}, Left: &TreeNode{Val: 4}}},
		out: 3,
	},
}

func TestDiameterOfBinary(t *testing.T) {
	for _, tt := range tts {
		got := diameterOfBinaryTree(tt.in)
		if got != tt.out {
			t.Errorf("Input: %+v | Expected: %d, Got: %d", tt.in, tt.out, got)
		}
	}
}
