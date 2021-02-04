/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// rl1 := reverse(l1)
	// rl2 := reverse(l2)

	rl1 := l1
	rl2 := l2

	sentinel := new(ListNode)
	cur := sentinel
	carry := 0

	for rl1 != nil && rl2 != nil {
		cur.Next = new(ListNode)
		cur = cur.Next
		cur.Val = (rl1.Val + rl2.Val + carry) % 10
		carry = (rl1.Val + rl2.Val + carry) / 10
		rl1 = rl1.Next
		rl2 = rl2.Next
	}

	for rl1 != nil {
		cur.Next = new(ListNode)
		cur = cur.Next
		cur.Val = (rl1.Val + carry) % 10
		carry = (rl1.Val + carry) / 10
		rl1 = rl1.Next
	}

	for rl2 != nil {
		cur.Next = new(ListNode)
		cur = cur.Next
		cur.Val = (rl2.Val + carry) % 10
		carry = (rl2.Val + carry) / 10
		rl2 = rl2.Next
	}

	if carry != 0 {
		cur.Next = new(ListNode)
		cur = cur.Next
		cur.Val = carry
	}

	return sentinel.Next

}

func reverse(root *ListNode) *ListNode {
	if root.Next == nil {
		return root
	}

	tail := reverse(root.Next)
	root.Next.Next = root
	root.Next = nil
	return tail
}

// @lc code=end

