package maxPathSum

import (
	"testing"
)

var tts = []struct {
	in  *TreeNode
	out int
}{
	{in: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, out: 6},
	{in: &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, out: 42},
	{in: &TreeNode{Val: -3}, out: -3},
}

func TestMaxPathSum(t *testing.T) {

	for _, tt := range tts {
		got := maxPathSum(tt.in)

		if tt.out != got {
			t.Errorf("Expected: %d; Got: %d", tt.out, got)
		}
	}
}
