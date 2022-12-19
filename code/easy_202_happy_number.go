/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

package code

// @lc code=start
//lint:ignore U1000 function is unused
func isHappy(n int) bool {
	// 方針
	// フロイドのサイクル発見アルゴリズムを使用する
	// スロウとファストを用意し、スロウは1回、ファストは2回ずつ進める
	// 2つが同じ値になったら、サイクルが発生している
	// ファストが1になったら、Happy Number(サイクルなし)

	// それぞれの位の2乗の和を計算する
	f := func(i int) int {
		sum := 0
		for i > 0 {
			sum += (i % 10) * (i % 10)
			i /= 10
		}
		return sum
	}

	slow, fast := n, f(n)

	// happy numberの場合、fastは1になるまでサイクルを繰り返す
	// happy numberではない場合、fast == slowになる
	for fast != 1 && slow != fast {
		slow = f(slow)
		fast = f(f(fast))
	}
	return fast == 1
}

// @lc code=end
