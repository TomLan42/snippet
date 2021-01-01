/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := new(ListNode)
	dummy.Next = head
	dummy.Val = 0

	current, pred := dummy, dummy
	counter := 0

	for counter < m {
		pred = current
		current = current.Next
		counter++
	}

	prev, front := current, current
	for counter <= n {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
		counter++
	}

	front.Next = current
	pred.Next = prev

	return dummy.Next
}

// @lc code=end

