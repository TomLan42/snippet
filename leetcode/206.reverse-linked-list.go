/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	//recursive way
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// @lc code=end

