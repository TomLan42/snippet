/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	Smalldummy := &ListNode{Val: 0}
	Bigdummy := &ListNode{Val: 0}

	Smallhead, Bighead := Smalldummy, Bigdummy

	for head != nil {
		if head.Val < x {
			Smallhead.Next = head
			Smallhead = Smallhead.Next
		} else {
			Bighead.Next = head
			Bighead = Bighead.Next
		}
		head = head.Next
	}

	Smallhead.Next = Bigdummy.Next
	Bighead.Next = nil
	return Smalldummy.Next

}

// @lc code=end

