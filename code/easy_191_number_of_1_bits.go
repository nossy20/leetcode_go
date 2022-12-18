/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

package code

// @lc code=start
//lint:ignore U1000 function is unused
func hammingWeight(num uint32) int {
	var res int = 0
	// 1ビット右シフトすることで、最下位ビットを取り出し、すべてのビットの1との論理積を取り、
	// その結果が1であれば、1の数をカウントアップする
	for i := 0; i < 32; i++ {
		tail := num & 1
		if tail == 1 {
			res++
		}
		// 1ビット右シフト
		num >>= 1
	}
	return res
}

// @lc code=end
