package l3

import "leec"

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	win := make(map[byte]int)
	l := 0
	for right < len(s) {
		c := s[right]
		win[c]++
		right++          // 窗口增大
		for win[c] > 1 { // 出现重复字符
			d := s[left]
			win[d]--
			left++
		}
		l = leec.Max(right-left, l)
	}
	return l
}
