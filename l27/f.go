package l27

func removeElement(nums []int, val int) int {
	low, fast := 0, len(nums)
	for low < fast {
		if nums[low] == val {
			nums[low] = nums[fast-1]
			fast--
		} else {
			low++
		}
	}
	return low
}

func removeElement2(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
