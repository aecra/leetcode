package offer28

import "testing"

// 剑指 Offer 28. 展开多级双向链表
func TestFlatten(t *testing.T) {
	// test case 1
	// 1 - 2 - 3
	//     |
	//     4 - 5 - 6
	//         |
	//         7
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}
	node1.Next = node2
	node2.Next = node3
	node2.Prev = node1
	node3.Prev = node2
	node2.Child = node4
	node4.Next = node5
	node5.Prev = node4
	node5.Next = node6
	node6.Prev = node5
	node5.Child = node7
	got := flatten(node1)
	resArrary := []int{1, 2, 4, 5, 7, 6, 3}
	for i := 1; i <= 7; i++ {
		if got.Val != resArrary[i-1] {
			t.Errorf("flatten() error.")
		}
		got = got.Next
	}
}
