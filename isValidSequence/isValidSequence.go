package isValidSequence

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidSequence(root *TreeNode, arr []int) bool {

	return helper(root, arr, 0)
}

func helper(root *TreeNode, arr []int, index int) bool {
	if root == nil || index == len(arr) {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == arr[index] && index == len(arr)-1 {
		return true
	}
	return root.Val == arr[index] && (helper(root.Left, arr, index+1) || helper(root.Right, arr, index+1))
}
