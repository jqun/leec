package l76

import "math"

func MinWindow(s string, t string) string {
	start, l := 0, math.MaxInt
	left, right := 0, 0

	need := make(map[byte]int) // 子串
	win := make(map[byte]int)  // 窗口
	valid := 0

	for _, v := range t {
		need[byte(v)]++
	}

	for right < len(s) { // 只需要遍历一次 O(N)
		c := s[right]
		right++ // 窗口增大
		// 更新数据
		if num, ok := need[c]; ok { // 更新操作A
			win[c]++ // 缓存
			if win[c] == num {
				valid++
			}
		}

		for valid == len(need) { // 所有字符找到后缩小窗口
			// 找到合法窗口则更新数据
			if right-left < l {
				start = left
				l = right - left
			}
			// 缩小窗口
			d := s[left]
			left++
			if num, ok := need[d]; ok { // 更新操作B，与操作A对称
				if win[d] == num {
					valid--
				}
				win[d]--
			}
		}
	}
	if l == math.MaxInt {
		return ""
	}
	return s[start : start+l]
}
