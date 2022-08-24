package problem1609

import "testing"

// 1609. Even Odd Tree
// 1609. 奇偶树
func TestEvenOddTree(t *testing.T) {
	// test case 1
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Left = &TreeNode{Val: 12}
	root.Left.Left.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}
	root.Right.Left.Left = &TreeNode{Val: 6}
	root.Right.Right.Right = &TreeNode{Val: 2}
	if isEvenOddTree(root) == false {
		t.Errorf("isEvenOddTree() error.")
	}
	// test case 2
	root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 7}
	if isEvenOddTree(root) == true {
		t.Errorf("isEvenOddTree() error.")
	}
	// test case 3
	root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 7}
	if isEvenOddTree(root) == true {
		t.Errorf("isEvenOddTree() error.")
	}
	root = &TreeNode{Val: 1}
	if isEvenOddTree(root) == false {
		t.Errorf("isEvenOddTree() error.")
	}
}
