package diameterOfBinaryTree

// TreeNode is data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var diameter int
	helper(root, &diameter)
	return diameter - 1

}

func helper(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, diameter)
	right := helper(root.Right, diameter)
	*diameter = max(*diameter, left+right+1)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
