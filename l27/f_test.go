package l27

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	x := removeElement(nums, 3)
	assert.Equal(t, 2, x)
}

func TestRemoveElement2(t *testing.T) {
	nums := []int{0,1,2,2,3,0,4,2}
	x := removeElement(nums, 2)
	assert.Equal(t, 5, x)
}
func TestRemoveElement3(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	x := removeElement2(nums, 3)
	assert.Equal(t, 2, x)
}

func TestRemoveElement4(t *testing.T) {
	nums := []int{0,1,2,2,3,0,4,2}
	x := removeElement2(nums, 2)
	assert.Equal(t, 5, x)
}