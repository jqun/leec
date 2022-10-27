package l438

// 滑动窗口
func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	need := make(map[byte]int)
	for _, v := range p {
		need[byte(v)]++
	}

	win := make(map[byte]int)
	valid := 0

	var res []int
	for right < len(s) {
		c := s[right]
		right++ // 扩大窗口
		if num, ok := need[c]; ok {
			win[c]++
			if win[c] == num { // 字符数量相等
				valid++
			}
		}
		for right-left >= len(p) { // 减小窗口
			if valid == len(need) { // 合法
				res = append(res, left)
			}
			d := s[left]
			left++
			if num, ok := need[d]; ok {
				if num == win[d] {
					valid--
				}
				win[d]--
			}
		}
	}
	return res
}
