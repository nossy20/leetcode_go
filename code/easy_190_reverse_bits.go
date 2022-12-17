/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func reverseBits(num uint32) uint32 {
	// 初期化
	var res uint32 = 0
	// 32ビットすべてを評価する
	for i := 0; i < 32; i++ {
		// 最下位bitが0なら0, 1なら1
		bom := num & 1
		// 最上位に挿入 -> 最下位に挿入したものを上位に繰り上げを(回数-1回分)繰り返す
		// 回数回繰り返してしまうと、シフトしすぎ
		// res<<1 ではi=0ではres=0なのでシフトしても無視される
		// i>0からシフトが意味を成すので、シフト回数はi(max)-1 = 31回で良し
		res = res<<1 | bom
		// 評価する数字を右に1ビットシフトして、次のビットを評価する
		num >>= 1
	}
	return res
}

// @lc code=end
