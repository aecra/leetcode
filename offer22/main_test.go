package offer22

import (
	"reflect"
	"testing"
)

// 剑指 Offer 22. 链表中倒数第k个节点
func TestGetKthFromEnd(t *testing.T) {
	// test case 1
	listofVal := []int{1, 2, 3, 4, 5}
	var list *ListNode
	for i := len(listofVal) - 1; i >= 0; i-- {
		list = &ListNode{listofVal[i], list}
	}
	got := getKthFromEnd(list, 2)
	res := []int{}
	for got != nil {
		res = append(res, got.Val)
		got = got.Next
	}
	if reflect.DeepEqual(res, []int{4, 5}) == false {
		t.Errorf("getKthFromEnd() error.")
	}
}
