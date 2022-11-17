/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

package code

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
//lint:ignore U1000 for leetcode
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// edge cases
	if list1 == nil && list2 != nil {
		return list2
	}
	// edge cases
	if list1 != nil && list2 == nil {
		return list1
	}

	// edge cases
	if list1 == nil && list2 == nil {
		return nil
	}

	// 再帰的に処理をする
	out := &ListNode{}
	if list1.Val <= list2.Val {
		out.Val = list1.Val
		list1 = list1.Next
	} else {
		out.Val = list2.Val
		list2 = list2.Next
	}
	out.Next = mergeTwoLists(list1, list2)
	return out
}

// @lc code=end
