package problem61

import "testing"

func TestRotateRight(t *testing.T) {
	// test case 1
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	resArrary := []int{4, 5, 1, 2, 3}
	got := rotateRight(node1, 2)
	for i := 0; i < len(resArrary); i++ {
		if got.Val != resArrary[i] {
			t.Errorf("rotateRight(%v, %d) = %d, want %d", node1, 2, got.Val, resArrary[i])
		}
		got = got.Next
	}
}
