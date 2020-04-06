package addTwoNumbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	l3 := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}

	output := addTwoNumbers(l1, l2)
	var expected, got []int

	for l := l3; l != nil; l = l.Next {
		expected = append(expected, l.Val)
	}

	for o := output; o != nil; o = o.Next {
		got = append(got, o.Val)
	}

	if !reflect.DeepEqual(output, got) {
		t.Errorf("Expected: %v; Got: %v", expected, got)
	}

}
