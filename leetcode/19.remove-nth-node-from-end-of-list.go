/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	sentinel := &ListNode{
		Next: head,
	}

	prevSlowNode := sentinel
	slowNode := head
	fastNode := head

	for fastNode != nil {
		if n <= 0 {
			prevSlowNode = slowNode
			slowNode = slowNode.Next
		}

		fastNode = fastNode.Next
		n--
	}

	prevSlowNode.Next = slowNode.Next

	return sentinel.Next
}

// @lc code=end

