/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func removeElements(head *ListNode, val int) *ListNode {
	// remove head
	for head != nil && head.Val == val {
		head = head.Next
	}

	// return case
	if head == nil {
		return nil
	}

	// recursive
	head.Next = removeElements(head.Next, val)
	return head
}

// @lc code=end
