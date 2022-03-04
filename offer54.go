package leetcode

import "context"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func work(ctx context.Context, root *TreeNode, ch chan<- int) {
	if root == nil {
		return
	}
	work(ctx, root.Right, ch)
	select {
	case ch <- root.Val:
	case <-ctx.Done():
		return
	}
	work(ctx, root.Left, ch)
}

func kthLargest(root *TreeNode, k int) int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan int)
	go work(ctx, root, ch)
	for i := 0; i < k-1; i++ {
		<-ch
	}
	return <-ch
}
