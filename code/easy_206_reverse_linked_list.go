/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
package code

// @lc code=start
//lint:ignore U1000 leetcode
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	// ポインタを手前を向くように差し替えていく
	// 1 -> 2 -> 3 -> 4 -> 5
	// 1 <- 2 <- 3 <- 4 <- 5
	// cur.Nextがprevを向くように更新していく
	// 差し替える前にcur.Nextを保存しておく(next)
	for cur != nil {
		next := cur.Next // 退避
		cur.Next = prev  // 差し替え
		prev = cur       // prevを1つ進める
		cur = next       // curを1つ進める(退避しておいたnext)
	}
	return prev
}

// @lc code=end
