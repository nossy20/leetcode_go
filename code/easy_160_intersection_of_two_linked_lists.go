/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 交差していれば、交差点から最後までの長さは同じなので、長い方のリストを先に進めて、深さを揃える
	// 共通部分で進める
	cA := headA
	cB := headB
	for cA != nil && cB != nil {
		cA = cA.Next
		cB = cB.Next
	}
	// cAが長い場合は、cAとheadAを長い分だけ進める
	for cA != nil {
		headA = headA.Next
		cA = cA.Next
	}
	// cBが長い場合は、cBとheadBを長い分だけ進める
	for cB != nil {
		headB = headB.Next
		cB = cB.Next
	}
	//　headAとheadBを同じ速度で進めていき、交差点を探す(headA == headBの位置が交差点)
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	// 交差している場合は、headA == headB == 交差点、交差していない場合は、headA == nil
	return headA
}

// @lc code=end
