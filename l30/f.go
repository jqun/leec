package l30

func findSubstring(s string, words []string) []int {
	n := len(s)
	oneWord := len(words[0])
	wordNum := len(words)
	totalLen := wordNum * oneWord
	if totalLen > n { // 长度不够
		return nil
	}

	need := make(map[string]int)
	for _, word := range words {
		need[word]++
	}

	var res []int
	for i := 0; i < len(s)-totalLen+1; i++ {
		tmpStr := s[i : i+totalLen]
		m := make(map[string]int)
		valid := 0
		for j := 0; j < totalLen; j += oneWord {
			v := tmpStr[j : j+oneWord]
			m[v]++
			if m[v] == need[v] {
				valid++
			}
		}
		if valid == len(need) {
			res = append(res, i)
		}
	}
	return res
}
