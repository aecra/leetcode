package problem1206

import (
	"math/rand"
	"time"
)

type node struct {
	value int
	count int
	right *node
	down  *node
}

type Skiplist struct {
	head   *node
	rate   int        // 表示相邻两层链表的元素数量之比
	level  int        // 表示最大层数
	length int        // 表示底层节点个数
	size   int        // 表示所有层节点个数
	rand   *rand.Rand // 用于随机数
	stack  []*node    // stack保存遍历路径中每一层最右边的节点
}

func Constructor() Skiplist {
	skiplist := Skiplist{
		head:   &node{value: -1},
		rate:   4,
		level:  8,
		length: 0,
		size:   0,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
		stack:  []*node{},
	}
	temp := skiplist.head
	for i := 1; i < skiplist.level; i++ {
		temp.down = &node{value: -1}
		temp = temp.down
	}
	return skiplist
}

func (skiplist *Skiplist) Search(target int) bool {
	skiplist.stack = skiplist.stack[:0]
	temp := skiplist.head
	for {
		for temp.right != nil {
			if temp.right.value >= target {
				break
			}
			temp = temp.right
		}
		if temp.down == nil {
			break
		}
		skiplist.stack = append(skiplist.stack, temp)
		temp = temp.down
	}
	skiplist.stack = append(skiplist.stack, temp)
	return temp.right != nil && temp.right.value == target
}

func (skiplist *Skiplist) Add(num int) {
	if skiplist.Search(num) {
		temp := skiplist.stack[len(skiplist.stack)-1]
		temp.right.count++
		return
	}
	temp := skiplist.stack[len(skiplist.stack)-1]
	newNode := &node{value: num, count: 1}
	var other *node
	// 根据随机数，自底向上添加每层的新节点
	for {
		newNode.right = temp.right
		temp.right = newNode
		skiplist.size++
		if skiplist.rand.Intn(skiplist.rate) != 0 || len(skiplist.stack) == 0 {
			break
		}
		//若随机数为 0且还未到顶层，则向上层添加元素
		temp = skiplist.stack[len(skiplist.stack)-1]
		skiplist.stack = skiplist.stack[:len(skiplist.stack)-1]
		other = newNode
		newNode = &node{value: num, count: 1}
		newNode.down = other
	}
	skiplist.length++
}

func (skiplist *Skiplist) Erase(num int) bool {
	if !skiplist.Search(num) {
		return false
	}
	temp := skiplist.stack[len(skiplist.stack)-1]
	if temp.right.count > 1 {
		temp.right.count--
		return true
	}
	for {
		if temp.right == nil || temp.right.value != num {
			break
		}
		//从底层开始，依次删除每层的元素
		temp.right = temp.right.right
		skiplist.size--
		if len(skiplist.stack) == 0 { //到达顶层
			break
		}
		temp = skiplist.stack[len(skiplist.stack)-1]
		skiplist.stack = skiplist.stack[:len(skiplist.stack)-1]
	}
	skiplist.length--
	return true
}

func (skiplist *Skiplist) Print() {
	head := skiplist.head
	for head != nil {
		temp := head
		for temp != nil {
			print(temp.value, " ")
			temp = temp.right
		}
		println()
		head = head.down
	}
}
