/*
 * @lc app=leetcode id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 */

// @lc code=start
type MedianFinder struct {
	max *maxheap
	min *minheap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{max: &maxheap{},
		min: &minheap{},
	}
}

func (this *MedianFinder) AddNum(num int) {

	a := this.max
	b := this.min
	if len(*b) == 0 {
		heap.Push(b, num)
		return
	}
	if len(*a) == len(*b) {
		if num < (*a)[0] {
			heap.Push(b, (*a)[0])
			heap.Pop(a)
			heap.Push(a, num)
		} else {
			heap.Push(b, num)
		}
	} else {
		if num > (*b)[0] {
			heap.Push(a, (*b)[0])
			heap.Pop(b)
			heap.Push(b, num)
		} else {
			heap.Push(a, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	a := this.max
	b := this.min
	if len(*b) == 0 {
		return 0
	}
	if len(*a) == len(*b) {
		return (float64((*a)[0]) + float64((*b)[0])) / 2
	} else {
		return float64((*b)[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type maxheap []int
type minheap []int

func (h maxheap) Len() int {
	return len(h)
}

func (h maxheap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h maxheap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxheap) Pop() interface{} {
	a := *h
	res := a[len(a)-1]
	*h = a[:len(a)-1]
	return res
}

func (h *maxheap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func (h minheap) Len() int {
	return len(h)
}

func (h minheap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h minheap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Pop() interface{} {
	a := *h
	res := a[len(a)-1]
	*h = a[:len(a)-1]
	return res
}

func (h *minheap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

