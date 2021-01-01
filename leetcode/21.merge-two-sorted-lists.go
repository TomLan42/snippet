/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := new(ListNode)
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next

		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}

	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}

	return dummy.Next
}

// @lc code=end

