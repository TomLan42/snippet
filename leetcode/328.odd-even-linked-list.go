/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	oddFront := head
	evenFront := head.Next
	for evenFront != nil && evenFront.Next != nil {
		// swap the oddFront and the evenFront

		cur := evenFront.Next
		tempOddN := oddFront.Next
		oddFront.Next = cur

		tempCurN := cur.Next
		cur.Next = tempOddN
		evenFront.Next = tempCurN

		oddFront = oddFront.Next
		evenFront = evenFront.Next

	}

	return head
}

// @lc code=end

