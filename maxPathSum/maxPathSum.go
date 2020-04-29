package maxPathSum

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maximum := math.MinInt32
	helper(&maximum, root)
	return maximum
}

func helper(maximum *int, root *TreeNode) int {

	if root == nil {
		return 0
	}

	var left, right int

	if root.Left == nil {
		left = 0
	} else {
		left = max(helper(maximum, root.Left), 0)
	}

	if root.Right == nil {
		right = 0
	} else {
		right = max(helper(maximum, root.Right), 0)
	}

	*maximum = max(left+right+root.Val, *maximum)
	return root.Val + max(left, right)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
