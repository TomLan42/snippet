/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		if current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
			continue
		}
		current = current.Next
	}
	return head
}

// @lc code=end

