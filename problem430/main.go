package problem430

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

var nodeList430 []*Node

func flatten430(root *Node) *Node {
	if root == nil {
		return nil
	}

	nodeList430 = make([]*Node, 0)
	dfs430(root)
	flatten := nodeList430[0]
	nodeList430[0].Child = nil
	for i := 1; i < len(nodeList430); i++ {
		nodeList430[i-1].Next = nodeList430[i]
		nodeList430[i].Prev = nodeList430[i-1]
		nodeList430[i].Child = nil
	}

	return flatten
}

func dfs430(root *Node) {
	if root == nil {
		return
	}
	nodeList430 = append(nodeList430, root)
	if root.Child != nil {
		dfs430(root.Child)
	}
	if root.Next != nil {
		dfs430(root.Next)
	}
}
