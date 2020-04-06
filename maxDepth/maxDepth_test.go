package maxDepth

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {

	in := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	out := 3

	got := maxDepth(in)

	if out != got {
		t.Errorf("Expected: %d; Got: %d", out, got)
	}
}
