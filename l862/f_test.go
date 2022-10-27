package l862

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestSubarray(t *testing.T) {
	nums := []int{1}
	k := 1
	x := shortestSubarray(nums, k)
	assert.Equal(t, 1, x)
}

func TestShortestSubarray2(t *testing.T) {
	nums := []int{1, 2}
	k := 4
	x := shortestSubarray(nums, k)
	assert.Equal(t, -1, x)
}

func TestShortestSubarray3(t *testing.T) {
	nums := []int{2, -1, 2}
	k := 3
	x := shortestSubarray(nums, k)
	assert.Equal(t, 3, x)
}

func TestShortestSubarray4(t *testing.T) {
	nums := []int{84, -37, 32, 40, 95}
	k := 167
	x := shortestSubarray(nums, k)
	assert.Equal(t, 3, x)
}
