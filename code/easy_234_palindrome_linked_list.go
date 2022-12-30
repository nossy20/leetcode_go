/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func isPalindrome(head *ListNode) bool {
	// edge case
	if head == nil || head.Next == nil {
		return true
	}

	// 今の順序を保存する
	p := []int{}
	c_head := head
	for c_head != nil {
		p = append(p, c_head.Val)
		c_head = c_head.Next
	}

	// 反転して順序を保存する
	a := []int{}
	prev := &ListNode{}
	cur := head
	for cur != nil {
		next := cur.Next // 退避
		cur.Next = prev  // 差し替え
		prev = cur       // prevを1つ進める
		cur = next       // curを1つ進める(退避しておいたnext)
	}
	for prev != nil {
		a = append(a, prev.Val)
		prev = prev.Next
	}
	// aの末尾にはprev初期化時のVal0が入っているので末尾を削除する
	return isSame(p, a[:len(a)-1])
}

func isSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// @lc code=end
