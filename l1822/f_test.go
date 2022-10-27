package l1822

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArraySign(t *testing.T) {
	nums := []int{-1, -2, -3, -4, 3, 2, 1}
	x := arraySign(nums)
	assert.Equal(t, 1, x)
}

func TestArraySign2(t *testing.T) {
	nums := []int{1,5,0,2,-3}
	x := arraySign(nums)
	assert.Equal(t, 0, x)
}

func TestArraySign3(t *testing.T) {
	nums := []int{-1,1,-1,1,-1}
	x := arraySign(nums)
	assert.Equal(t, -1, x)
}

func TestArraySign4(t *testing.T) {
	nums := []int{41,65,14,80,20,10,55,58,24,56,28,86,96,10,3,84,4,41,13,32,42,43,83,78,82,70,15,-41}
	x := arraySign(nums)
	assert.Equal(t, -1, x)
}
