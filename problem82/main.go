package problem82

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

func deleteDuplicates(head *ListNode) *ListNode {
	m := make(map[int]int)
	p := head
	for {
		if p == nil {
			break
		}
		_, ok := m[p.Val]
		if ok {
			m[p.Val]++
		} else {
			m[p.Val] = 1
		}
		p = p.Next
	}

	list := []int{}
	for index := range m {
		if m[index] < 2 {
			list = append(list, index)
		}
	}
	InsertionSort(list)
	var res *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		res = &ListNode{list[i], res}
	}
	return res
}

func InsertionSort(A []int) {
	for i := 1; i < len(A); i++ {
		for j := i; j > 0 && A[j] < A[j-1]; j-- {
			A[j], A[j-1] = A[j-1], A[j]
		}
	}
}
