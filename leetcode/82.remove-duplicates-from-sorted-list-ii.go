/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
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

	sentinel := &ListNode{Val: 0, Next: head}
	lastDistinct := sentinel
	current := head

	for current != nil {
		if current.Next != nil && current.Val == current.Next.Val {
			for current.Next != nil && current.Val == current.Next.Val {
				current = current.Next
			}
			lastDistinct.Next = current.Next
		} else {
			lastDistinct = current
		}
		current = current.Next
	}
	return sentinel.Next

}

// essense of this problem is how to remove a contiguous part of
// of a single linked list properply

// @lc code=end

