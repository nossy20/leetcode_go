/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

package code

// @lc code=start
//lint:ignore U1000 for leetcode
func deleteDuplicates(head *ListNode) *ListNode {
	// edge cases
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	// もし、head.Next.Val が head.Val と同じなら、head.Next を head.Next.Next に置き換える
	if head.Next.Val == head.Val {
		head.Next = head.Next.Next
		return deleteDuplicates(head)
	}

	// 再帰的に処理をする
	head.Next = deleteDuplicates(head.Next)
	return head
}

// @lc code=end
