package l1822

func arraySign(nums []int) int {
	var mul = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 { // 不必走后续流程
			return 0
		} else if nums[i] < 0 {
			mul *= -1
		}
	}
	return mul
}
