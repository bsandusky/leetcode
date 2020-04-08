package middleNode

import (
	"testing"
)

var tts = []struct {
	in  *ListNode
	out *ListNode
}{
	{
		in:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
		out: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
	},
	{
		in:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}},
		out: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}},
	},
}

func TestMiddleNode(t *testing.T) {
	for _, tt := range tts {
		got := middleNode(tt.in)
		if got.Val != tt.out.Val {
			t.Errorf("Input: %+v | Expected: %+v, Got: %+v", tt.in.Val, tt.out.Val, got.Val)
		}
	}
}
