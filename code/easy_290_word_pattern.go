/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

package code

import (
	"strings"
)

// @lc code=start
//lint:ignore U1000 LeetCode
func wordPattern(pattern string, s string) bool {
	// patternをスライスに変換
	p := []string{}
	for _, c := range pattern {
		p = append(p, string(c))
	}
	// s文字列をスライスに変換
	words := strings.Split(s, " ")

	// 両方のスライスの長さが異なる場合はfalse
	if len(p) != len(words) {
		return false
	}

	// patternとs文字列の対応関係を保持するマップ
	m := map[string]string{}
	for i := 0; i < len(p); i++ {
		// patternの文字がマップに存在しない場合
		if _, ok := m[p[i]]; !ok {
			// s文字列の文字がマップに存在する場合はfalse
			for _, v := range m {
				if v == words[i] {
					return false
				}
			}
			// マップに追加
			m[p[i]] = words[i]
		} else {
			// patternの文字がマップに存在する場合
			// s文字列の文字がマップの値と異なる場合はfalse
			if m[p[i]] != words[i] {
				return false
			}
		}
	}

	return true
}

// @lc code=end
