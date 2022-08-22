package leetcode

import (
	"reflect"
	"testing"
)

// 8. String to Integer (atoi)
// 8. 字符串转换整数 (atoi)
func TestMyAtoi(t *testing.T) {
	// test case 1
	got := myAtoi("42")
	if got != 42 {
		t.Errorf("myAtoi() error: 42 -> %d", got)
	}

	// test case 2
	got = myAtoi("   -42")
	if got != -42 {
		t.Errorf("myAtoi() error: -42 -> %d", got)
	}

	// test case 3
	got = myAtoi("4193 with words")
	if got != 4193 {
		t.Errorf("myAtoi() error: 4193 -> %d", got)
	}
}

// 72. Edit Distance
// 72. 编辑距离
func TestMinDistance(t *testing.T) {
	// test case 1
	got := minDistance("horse", "ros")
	if got != 3 {
		t.Errorf("minDistance() error %d.", got)
	}

	// test case 2
	got = minDistance("intention", "execution")
	if got != 5 {
		t.Errorf("minDistance() error %d.", got)
	}
}

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

// 84. Largest Rectangle in Histogram
// 84. 柱状图中最大的矩形
func TestLargestRectangleArea(t *testing.T) {
	// test case 1
	got := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	if got != 10 {
		t.Errorf("largestRectangleArea() error. %d", got)
	}

	// test case 2
	got = largestRectangleArea([]int{4, 2, 0, 3, 2, 5})
	if got != 6 {
		t.Errorf("largestRectangleArea() error. %d", got)
	}

	// test case 3
	got = largestRectangleArea([]int{4})
	if got != 4 {
		t.Errorf("largestRectangleArea() error. %d", got)
	}
}

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

// 309. Best Time to Buy and Sell Stock with Cooldown
// 309. 买卖股票的最佳时机含冷冻期
func TestMaxProfit309(t *testing.T) {
	// test case 1
	got := maxProfit309([]int{1, 2, 3, 0, 2})
	if got != 3 {
		t.Errorf("maxProfit309() error %d.", got)
	}

	// test case 2
	got = maxProfit309([]int{1})
	if got != 0 {
		t.Errorf("maxProfit309() error %d.", got)
	}
}

// 343. Integer Break
// 343. 整数拆分
func TestIntegerBreak(t *testing.T) {
	// test case 1
	got := integerBreak(2)
	if got != 1 {
		t.Errorf("integerBreak() error %d.", got)
	}

	// test case 2
	got = integerBreak(3)
	if got != 2 {
		t.Errorf("integerBreak() error %d.", got)
	}

	// test case 3
	got = integerBreak(9)
	if got != 27 {
		t.Errorf("integerBreak() error %d.", got)
	}

	// test case 4
	got = integerBreak(10)
	if got != 36 {
		t.Errorf("integerBreak() error %d.", got)
	}
}

// 402. Remove K Digits
// 402. 移掉 K 位数字
func TestRemoveKdigits(t *testing.T) {
	// test case 1
	got := removeKdigits("1432219", 3)
	if got != "1219" {
		t.Errorf("removeKdigits() error %s.", got)
	}

	// test case 2
	got = removeKdigits("10200", 1)
	if got != "200" {
		t.Errorf("removeKdigits() error %s.", got)
	}

	// test case 3
	got = removeKdigits("10", 1)
	if got != "0" {
		t.Errorf("removeKdigits() error %s.", got)
	}

	// test case 4
	got = removeKdigits("1140004278", 4)
	if got != "278" {
		t.Errorf("removeKdigits() error %s.", got)
	}

	// test case 5
	got = removeKdigits("278", 2)
	if got != "2" {
		t.Errorf("removeKdigits() error %s.", got)
	}
}

// 557. Reverse Words in a String III
// 557. 反转字符串中的单词 III
func TestReverseWords(t *testing.T) {
	// test case 1
	got := reverseWords("Let's take LeetCode contest")
	if got != "s'teL ekat edoCteeL tsetnoc" {
		t.Errorf("reverseWords() error %s.", got)
	}

	// test case 2
	got = reverseWords("God Ding")
	if got != "doG gniD" {
		t.Errorf("reverseWords() error %s.", got)
	}
}

// 658. Find K Closest Elements
// 658. 找到 K 个最接近的元素
func TestFindClosestElements(t *testing.T) {
	// test case 1
	got := findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3)
	if reflect.DeepEqual(got, []int{1, 2, 3, 4}) == false {
		t.Errorf("findClosestElements() error.")
	}

	// test case 2
	got = findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1)
	if reflect.DeepEqual(got, []int{1, 2, 3, 4}) == false {
		t.Errorf("findClosestElements() error.")
	}
}

// 713. Subarray Product Less Than K
// 713. 乘积小于 K 的子数组
func TestNumSubarrayProductLessThanK(t *testing.T) {
	// test case 1
	got := numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)
	if got != 8 {
		t.Errorf("numSubarrayProductLessThanK() error %d.", got)
	}
}

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

// 1648. Maximum Product of Two Elements in an Array
// 1648. 销售价值减少的颜色球
func TestMaxProfit(t *testing.T) {
	// test case 1
	got := maxProfit([]int{2, 5}, 4)
	if got != 14 {
		t.Errorf("maxProfit() error: 14 -> %d", got)
	}

	// test case 2
	got = maxProfit([]int{1000000000}, 1000000000)
	if got != 21 {
		t.Errorf("maxProfit() error: 21 -> %d", got)
	}

	// test case 3
	got = maxProfit([]int{2, 8, 4, 10, 6}, 30)
	if got != 125 {
		t.Errorf("maxProfit() error: 125 -> %d", got)
	}

	// test case 4
	got = maxProfit([]int{680754224, 895956841, 524658639, 3161416, 992716630, 7365440, 582714485, 422256708, 332815744, 269407767}, 707568720)
	if got != 6150955 {
		t.Errorf("maxProfit() error: 6150955 -> %d", got)
	}
}

// 1827. Minimum Operations to Make the Array Increasing
// 1827. 最少操作使数组递增
func TestMinOperations(t *testing.T) {
	// test case 1
	got := minOperations([]int{1, 1, 1})
	if got != 3 {
		t.Errorf("minOperations() error %d.", got)
	}

	// test case 2
	got = minOperations([]int{1, 5, 2, 4, 1})
	if got != 14 {
		t.Errorf("minOperations() error %d.", got)
	}
}

// 2178. Maximum Split of Positive Even Integers
// 2178. 拆分成最多数目的正偶数之和
func TestMaximumEvenSplit(t *testing.T) {
	// test case 1
	got := maximumEvenSplit(12)
	if reflect.DeepEqual(got, []int64{2, 4, 6}) == false {
		t.Errorf("maximumEvenSplit() error.")
	}

	// test case 2
	got = maximumEvenSplit(7)
	if len(got) != 0 {
		t.Errorf("maximumEvenSplit() error.")
	}

	// test case 3
	got = maximumEvenSplit(28)
	if reflect.DeepEqual(got, []int64{2, 4, 6, 16}) == false {
		t.Errorf("maximumEvenSplit() error.")
	}
}

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
