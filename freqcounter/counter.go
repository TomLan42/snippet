package main

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

// Counter desctribed above
type Counter struct {
	NumTable map[int]int
}

// Add takes in number
func (c *Counter) Add(num int) {
	c.NumTable[num]++

}

// LastNonRepeat reports last non repeated number
func (c *Counter) LastNonRepeat() {

}

// MostFrequent returns the most frequent number observed
func (c *Counter) MostFrequent() {

}

// LeastFrequent returns the least frequent number observed
func (c *Counter) LeastFrequent() {

}

func main() {
	counter := Counter{
		NumTable: make(map[int]int),
	}

	counter.Add(1)
	counter.Add(2)
	// Should print 2
	counter.LastNonRepeat()
	counter.Add(2)
	// Should print 1
	counter.LastNonRepeat()
	counter.Add(2)
	counter.Add(3)
	// Should print 3
	counter.LastNonRepeat()
	counter.Add(3)
	counter.Add(3)
	counter.Add(3)
	counter.Add(3)
	// Should print 3
	counter.MostFrequent()
	// Should print 1
	counter.LeastFrequent()

}
