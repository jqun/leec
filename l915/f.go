package l915

import (
	"leec"
)

func partitionDisjoint(nums []int) int {
	n := len(nums)
	leftMax, leftPos, curMax := nums[0], 0, nums[0]
	for i := 1; i < n-1; i++ {
		curMax = leec.Max(nums[i], curMax)
		if nums[i] < leftMax {
			leftMax = curMax
			leftPos = i
		}
	}
	return leftPos + 1
}
