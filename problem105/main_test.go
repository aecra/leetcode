package problem105

import "testing"

// 105. Construct Binary Tree from Preorder and Inorder Traversal
// 105. 从先序与中序遍历序列构造二叉树
func TestBuildTree(t *testing.T) {
	// test case 1
	got := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	if got == nil {
		t.Errorf("buildTree() error. nil")
	}

	// test case 2
	got = buildTree([]int{1, 2, 3}, []int{2, 1, 3})
	if got == nil {
		t.Errorf("buildTree() error. nil")
	}
}
