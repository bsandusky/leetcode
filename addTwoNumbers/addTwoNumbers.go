package addTwoNumbers

// ListNode is the definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	result := &ListNode{}
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {

		sum := 0
		result.Next = &ListNode{}
		result = result.Next
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		if carry != 0 {
			sum = sum + carry
		}

		if sum > 9 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		result.Val = sum
	}
	return result.Next
}
