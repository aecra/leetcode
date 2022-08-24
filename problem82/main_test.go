package problem82

import (
	"reflect"
	"testing"
)

// 82. Remove Duplicates from Sorted List II
// 82. 删除排序链表中的重复元素 II
func TestDeleteDuplicates(t *testing.T) {
	// test case 1
	listofVal := []int{1, 2, 3, 3, 4, 4, 5}
	var list *ListNode
	for i := len(listofVal) - 1; i >= 0; i-- {
		list = &ListNode{listofVal[i], list}
	}
	got := deleteDuplicates(list)
	res := []int{}
	for got != nil {
		res = append(res, got.Val)
		got = got.Next
	}
	if reflect.DeepEqual(res, []int{1, 2, 5}) == false {
		t.Errorf("deleteDuplicates() error.")
	}

	// test case 2
	listofVal = []int{1, 1, 1, 2, 3}
	list = nil
	for i := len(listofVal) - 1; i >= 0; i-- {
		list = &ListNode{listofVal[i], list}
	}
	got = deleteDuplicates(list)
	res = []int{}
	for got != nil {
		res = append(res, got.Val)
		got = got.Next
	}
	if reflect.DeepEqual(res, []int{2, 3}) == false {
		t.Errorf("deleteDuplicates() error.")
	}
}
