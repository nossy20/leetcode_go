/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 */

package code

// @lc code=start
//lint:ignore U1000 used in leetcode
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 一方がnil, 一方が値が存在の場合はfalse
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		// LeftとRightを再帰的に比較する。両方nilになるまで等しければtrueが返る
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	// 両方nilの場合はtrue
	return true
}

// @lc code=end
