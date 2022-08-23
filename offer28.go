package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

var nodeList []*Node

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	nodeList = make([]*Node, 0)
	dfs(root)
	flatten := nodeList[0]
	nodeList[0].Child = nil
	for i := 1; i < len(nodeList); i++ {
		nodeList[i-1].Next = nodeList[i]
		nodeList[i].Prev = nodeList[i-1]
		nodeList[i].Child = nil
	}

	return flatten
}

func dfs(root *Node) {
	if root == nil {
		return
	}
	nodeList = append(nodeList, root)
	if root.Child != nil {
		dfs(root.Child)
	}
	if root.Next != nil {
		dfs(root.Next)
	}
}
