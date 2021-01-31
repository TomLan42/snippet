/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	return helper(head)
}

func helper(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	temp := head.Next
	head.Next = helper(head.Next.Next)
	temp.Next = head
	return temp
}

// @lc code=end

