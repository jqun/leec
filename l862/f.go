package l862

import "leec"

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSumArr := make([]int, n+1)
	for i, num := range nums {
		preSumArr[i+1] = preSumArr[i] + num // 前缀和
	}
	var pre []int // 遍历前缀
	l := n + 1
	for i, sum := range preSumArr {
		for len(pre) > 0 && sum-preSumArr[pre[0]] >= k { // 找到目标-类似窗口缩小
			l = leec.Min(l, i-pre[0])
			pre = pre[1:]
		}
		for len(pre) > 0 && preSumArr[pre[len(pre)-1]] >= sum { // 单调递增序列
			pre = pre[:len(pre)-1]
		}
		pre = append(pre, i)
	}
	if l == n+1 {
		return -1
	}
	return l
}
