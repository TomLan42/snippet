/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

//First method: hashmap to store address. trivial
//Second method: put clone next to original, then split

func copyRandomList(head *Node) *Node {

	if head == nil {
		return head
	}

	//clone
	current := head
	for current != nil {
		clone := &Node{Val: current.Val, Next: current.Next, Random: current.Random}
		temp := current.Next
		current.Next = clone
		current = temp
	}

	//shift
	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current = current.Next.Next
	}

	//split
	current = head
	cloneCurrent := current.Next

	for current != nil && current.Next != nil {
		temp := current.Next
		current.Next = current.Next.Next
		current = temp
	}

	return cloneCurrent

}

// @lc code=end

