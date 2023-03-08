package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }     // 大顶堆，返回值决定是否交换元素
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	pq := &IntHeap{}
	heap.Init(pq)
	heap.Push(pq, 3)
	heap.Push(pq, 4)
	heap.Push(pq, 2)
	heap.Push(pq, 5)
	heap.Push(pq, 1)

	//从小到大pop
	for len(*pq) > 0 {
		x := heap.Pop(pq)
		fmt.Println(x)
	}
}
