package offer54

import "testing"

// 剑指 Offer 54. 二叉搜索树的第k大节点
func TestKthLargest(t *testing.T) {
	// test case 1
	var n1, n2, n3, n4 TreeNode
	n1.Val = 3
	n2.Val = 1
	n3.Val = 2
	n4.Val = 4
	n1.Left = &n2
	n1.Right = &n4
	n2.Right = &n3
	got := kthLargest(&n1, 1)
	if got != 4 {
		t.Errorf("kthLargest() error.")
	}

	// test case 2
	var m1, m2, m3, m4, m5, m6 TreeNode
	m1.Val = 5
	m2.Val = 3
	m3.Val = 2
	m4.Val = 1
	m5.Val = 4
	m6.Val = 6
	m1.Left = &m2
	m1.Right = &m6
	m2.Left = &m3
	m2.Right = &m5
	m3.Left = &m4
	got = kthLargest(&m1, 3)
	if got != 4 {
		t.Errorf("kthLargest() error.")
	}
}
