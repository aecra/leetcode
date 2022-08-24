package problem61

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}

 func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
			return head
	}
	num := numOfList(head)
	offset := k % num
	if offset == 0 {
			return head
	}
	start := head
	var end *ListNode
	for i := 0; i<num-offset; i++ {
			end = start
			start = start.Next
	}
	end.Next = nil
	endNodeOfList(start).Next = head
	return start
}

func numOfList(head *ListNode) (num int) {
	for head != nil {
			num++;
			head = head.Next
	}
	return num
}

func endNodeOfList(head *ListNode) *ListNode {
	for head.Next != nil {
			head = head.Next
	}
	return head
}