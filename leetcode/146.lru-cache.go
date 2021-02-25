/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type LRUCache struct {
	Capacity int
	Size     int
	Head     *Node
	Tail     *Node
	Hashmap  map[int]*Node
}

type Node struct {
	Key      int
	Val      int
	Next     *Node
	Previous *Node
}

func Constructor(capacity int) LRUCache {

	head := new(Node)
	tail := new(Node)
	head.Next = tail
	tail.Previous = head

	return LRUCache{
		Capacity: capacity,
		Head:     head,
		Tail:     tail,
		Hashmap:  make(map[int]*Node),
	}

}

func (this *LRUCache) Get(key int) int {

	// move the node to the front
	if target, found := this.Hashmap[key]; !found {
		return -1
	} else {

		target.Previous.Next = target.Next
		target.Next.Previous = target.Previous

		this.Head.Next.Previous = target
		target.Next = this.Head.Next
		target.Previous = this.Head
		this.Head.Next = target

	}
	// get the value
	return this.Head.Next.Val
}

func (this *LRUCache) Put(key int, value int) {

	var node *Node
	if target, found := this.Hashmap[key]; !found {
		// insert the node to the front
		node = &Node{
			Key: key,
			Val: value,
		}

		this.Size++
		this.Hashmap[key] = node
		// evict the tail if cap exceeds
		if this.Size > this.Capacity {
			delete(this.Hashmap, this.Tail.Previous.Key)
			this.Tail.Previous.Previous.Next = this.Tail
			this.Tail.Previous = this.Tail.Previous.Previous
			this.Size--
		}

	} else {
		node = target
		node.Val = value

		node.Previous.Next = node.Next
		node.Next.Previous = node.Previous
	}

	// move to head

	this.Head.Next.Previous = node
	node.Next = this.Head.Next
	node.Previous = this.Head
	this.Head.Next = node

}

func (this *LRUCache) check() {

	cur := this.Head

	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("\n")
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

