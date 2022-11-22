/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

package code

// @lc code=start
//lint:ignore U1000 func mySqrt(x int) int {
func mySqrt(x int) int {
	for i := 1; i <= x; i++ {
		if i*i == x {
			return i
		}
		// 2乗した値がxを超えたら、i-1を返す
		// x = 8, 3 * 3 = 9 > 8 なので、3-1 = 2を返す
		if i*i > x {
			return i - 1
		}
	}
	return 0
}

// @lc code=end
