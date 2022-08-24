package offer22

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	num := 0
	start := head
	for start != nil {
		num++
		start = start.Next
	}
	for i := 0; i < num-k; i++ {
		head = head.Next
	}
	return head
}
