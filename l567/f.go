package l567

func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	need := make(map[byte]int)
	win := make(map[byte]int)
	for _, v := range s1 {
		need[byte(v)]++
	}
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if num, ok := need[c]; ok {
			win[c]++
			if win[c] == num {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if num, ok := need[d]; ok {
				if num == win[d] {
					valid--
				}
				win[d]--
			}
		}
	}
	return false
}
