package middleNode

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	length := size(head)
	middle := length / 2
	return getNodeAtIndex(head, middle)
}

func size(node *ListNode) int {
	size := 1
	last := node
	for {
		if last == nil || last.Next == nil {
			break
		}
		last = last.Next
		size++
	}
	return size
}

func getNodeAtIndex(node *ListNode, index int) *ListNode {

	current := node
	var count int

	for current != nil {
		if count == index {
			return current
		}
		count++
		current = current.Next
	}
	return nil
}
