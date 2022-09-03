package problem2102

import (
	"container/heap"
)

type SORTracker struct {
	better *PriorityQueue
	worse  *PriorityQueue
}

func Constructor() SORTracker {
	return SORTracker{&PriorityQueue{isBetter: true}, &PriorityQueue{isBetter: false}}
}

func (sorTracker *SORTracker) Add(name string, score int) {
	sorTracker.better.push(Sight{name, score})
	sorTracker.worse.push(sorTracker.better.pop())
}

func (sorTracker *SORTracker) Get() string {
	p := sorTracker.worse.pop()
	sorTracker.better.push(p)
	return p.name
}

type Sight struct {
	name  string
	score int
}

type PriorityQueue struct {
	data     []Sight
	isBetter bool
}

func (pq PriorityQueue) Len() int { return len(pq.data) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq.isBetter {
		return pq.data[i].score < pq.data[j].score || (pq.data[i].score == pq.data[j].score && pq.data[i].name > pq.data[j].name)
	} else {
		return pq.data[i].score > pq.data[j].score || (pq.data[i].score == pq.data[j].score && pq.data[i].name < pq.data[j].name)
	}
}
func (pq PriorityQueue) Swap(i, j int)       { pq.data[i], pq.data[j] = pq.data[j], pq.data[i] }
func (pq *PriorityQueue) Push(v interface{}) { pq.data = append(pq.data, v.(Sight)) }
func (pq *PriorityQueue) Pop() (v interface{}) {
	pq.data, v = pq.data[:len(pq.data)-1], pq.data[len(pq.data)-1]
	return
}
func (pq *PriorityQueue) push(v Sight) { heap.Push(pq, v) }
func (pq *PriorityQueue) pop() Sight   { return heap.Pop(pq).(Sight) }
