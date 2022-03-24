package problems

import (
	"math"
	"sort"
)

/*
//solution with Priority Queue
type IntHeap []int

func (t IntHeap) Len() int { return len(t) }
func (t IntHeap) Less(i, j int) bool {
	a := math.Abs(float64(t[i] - arrM))
	b := math.Abs(float64(t[j] - arrM))
	if a > b || a == b && t[i] > t[j] {
		return false
	}
	return true
}
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
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

var arrM int

func getStrongest(arr []int, k int) []int {

	sort.Ints(arr)
	arrM = arr[(len(arr)-1)/2]
	//Priority Queue
	topk := IntHeap{}
	h := &topk
	heap.Init(h)
	n := len(arr)
	for i := 0; i < n; i++ {
		if i < k {
			heap.Push(h, arr[i])
			continue
		}
		a := math.Abs(float64(arr[i] - arrM))
		b := math.Abs(float64(topk[0] - arrM))
		if a > b || a == b && arr[i] > topk[0] {
			heap.Pop(h)
			heap.Push(h, arr[i])
		}

	}
	return topk
}
*/
// solution use double ptr
func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	arrM := arr[(len(arr)-1)/2]
	i := 0
	j := len(arr) - 1
	h := []int{}
	for i <= j && len(h) < k {

		a := math.Abs(float64(arr[i] - arrM))
		b := math.Abs(float64(arr[j] - arrM))
		if i == j || a > b || a == b && arr[i] > arr[j] {
			h = append(h, arr[i])
			i++
		} else {
			h = append(h, arr[j])
			j--
		}
	}
	return h
}
