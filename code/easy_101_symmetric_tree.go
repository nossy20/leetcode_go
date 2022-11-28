/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 */
package code

// @lc code=start
//lint:ignore U1000 leetcode
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// ルートは再帰処理から除外する
	return equal(root.Left, root.Right)
}

func equal(l, r *TreeNode) bool {
	// 両方ともnilの場合はtrueを返却
	if l == nil && r == nil {
		return true
	}
	// どちらかがnilの場合はfalseを返却
	if l == nil && r != nil || l != nil && r == nil {
		return false
	}

	// いずれもnilになるまで再帰処理を行い、全てがtrueの場合はシンメトリックなツリーと判定する
	// シンメトリックの条件は、左の左と右の右、左の右と右の左が同じ値であること
	return l.Val == r.Val && equal(l.Left, r.Right) && equal(l.Right, r.Left)
}

// @lc code=end
