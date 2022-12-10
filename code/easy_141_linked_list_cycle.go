/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

package code

// @lc code=start
//lint:ignore U1000 leetcode
func hasCycle(head *ListNode) bool {
	// headがnilの場合はループが無いのでfalseを返す
	if head == nil {
		return false
	}

	// headがnilでない場合は、fast, slowでずらしを変えてループを回し、ループを見つける
	fast := head.Next
	slow := head
	for fast != slow && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// nil系でforが終了した場合はfast, slowが異なり、それ以外の終了はfast, slowが同じになるのでループが存在する
	return fast == slow
}

// @lc code=end
