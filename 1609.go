package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	floorElement := make(map[int]int)
	return InorderTraversal(root, 0, floorElement)
}

func InorderTraversal(node *TreeNode, depth int, floorElement map[int]int) bool {
	if depth%2 == 0 && node.Val%2 == 0 {
		return false
	}
	if depth%2 == 1 && node.Val%2 == 1 {
		return false
	}
	floorVal, ok := floorElement[depth]
	floorElement[depth] = node.Val
	if ok {
		if depth%2 == 0 && node.Val <= floorVal {
			return false
		}
		if depth%2 == 1 && node.Val >= floorVal {
			return false
		}
	}
	if node.Left != nil && node.Right != nil {
		if depth%2 == 0 && node.Left.Val <= node.Right.Val {
			return false
		}
		if depth%2 == 1 && node.Left.Val >= node.Right.Val {
			return false
		}
	}
	if node.Left != nil {
		if !InorderTraversal(node.Left, depth+1, floorElement) {
			return false
		}
	}
	if node.Right != nil {
		if !InorderTraversal(node.Right, depth+1, floorElement) {
			return false
		}
	}
	return true
}
