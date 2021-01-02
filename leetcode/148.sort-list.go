/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	result := mergeSort(head)
	return result
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middle := findMiddle(head)
	tail := middle.Next
	middle.Next = nil

	left := mergeSort(head)
	right := mergeSort(tail)

	return mergeTwoLists(left, right)

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{Next: nil, Val: 0}
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

func findMiddle(head *ListNode) *ListNode {
	// very tricky here, because we will split the list
	// from the next node after the middle this function return
	// so this functin must return a 'middle' of [n-1]/2
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow

}

// @lc code=end

