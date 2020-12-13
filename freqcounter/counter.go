package main

import "fmt"

/*
 * @Author: TomLan42
 * @Date: 2020-12-10 23:19:01
 * @Last Modified by: TomLan42
 * @Last Modified time: 2020-12-10 23:56:13
 */

// Design a class that:
// 1. takes in number and reports last non repeated number.
// 2. report most/least frequent number.
// 3. mult-thread scenario.

// Counter abstract
type Counter interface {
	Add(int)
	LastNonRepeat()
	MostFrequent()
	LeastFrequent()
}

// NaiveCounter uses map
type NaiveCounter struct {
	Sequence  []int
	Frequency map[int]int
}

// LinearCounter uses map + linear insertion
type LinearCounter struct {
}

// HeapCounter uses map + heap
type HeapCounter struct {
}

// Add takes in number
func (n *NaiveCounter) Add(num int) {
	n.Sequence = append(n.Sequence, num)
	n.Frequency[num]++
	fmt.Println(n.Sequence)
}

// LastNonRepeat method of NaiveCounter
// Time complexity: O(n)
func (n *NaiveCounter) LastNonRepeat() {
	ans := -1
	for _, num := range n.Sequence {
		if n.Frequency[num] == 1 {
			ans = num
		}
	}
	fmt.Println("Last non-repeated is: ", ans)
}

// MostFrequent method of NaiveCounter
func (n *NaiveCounter) MostFrequent() {
	ans := make([]int, 0)
	max := 0
	for _, freq := range n.Frequency {
		if freq >= max {
			max = freq
		}
	}
	for num, freq := range n.Frequency {
		if freq == max {
			ans = append(ans, num)
		}
	}
	fmt.Println("Most frequent is (are): ", ans)

}

// LeastFrequent method of NaiveCounter
func (n *NaiveCounter) LeastFrequent() {

}

func main() {
	counterA := &NaiveCounter{
		Sequence:  make([]int, 0),
		Frequency: make(map[int]int),
	}
	// counterB := HeapCounter{}
	fmt.Println("*** Naive Counter ***")
	Test(counterA)
	// fmt.Println("*** Heap Counter ***")
	// Test(counterB)

}

// Test function test the coutner
func Test(ctr Counter) {
	ctr.Add(1)
	ctr.Add(2)
	// Should print 2
	ctr.LastNonRepeat()
	ctr.Add(2)
	// Should print 1
	ctr.LastNonRepeat()
	ctr.Add(2)
	ctr.Add(3)
	// Should print 3
	ctr.LastNonRepeat()
	ctr.Add(3)
	ctr.Add(3)
	ctr.Add(3)
	ctr.Add(3)
	// Should print 3
	ctr.MostFrequent()
	// Should print 1
	ctr.LeastFrequent()

}
