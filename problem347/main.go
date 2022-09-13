package problem347

import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	frequence := map[int]int{}
	for _, num := range nums {
		frequence[num]++
	}
	pq := make(PriorityQueue, 0, k)
	for num, count := range frequence {
		heap.Push(&pq, &Item{num, count})
		if len(pq) > k {
			heap.Pop(&pq)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(&pq).(*Item).value
	}
	return res
}

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}
