package problem1354

import (
	"container/heap"
)

// implement a priority queue
type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
func (pq *PriorityQueue) Top() int {
	return (*pq)[0]
}

func isPossible(target []int) bool {
	if len(target) == 1 {
		return target[0] == 1
	}
	var sum int64
	pq := make(PriorityQueue, 0)
	for _, v := range target {
		sum += int64(v)
		pq.Push(v)
	}
	heap.Init(&pq)

	for sum > int64(len(target)) {
		maxValue := heap.Pop(&pq).(int)
		secondMaxValue := pq.Top()
		rest := sum - int64(maxValue)
		var n int64
		if maxValue == secondMaxValue || int64(maxValue-secondMaxValue)%(rest) != 0 {
			n = int64(maxValue-secondMaxValue)/rest + 1
		} else {
			n = int64(maxValue-secondMaxValue) / rest
		}
		newValue := maxValue - int(n*(sum-int64(maxValue)))

		if newValue <= 0 {
			return false
		}
		heap.Push(&pq, newValue)
		sum -= int64(maxValue - newValue)
	}
	return sum == int64(len(target))
}
