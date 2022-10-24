package l915

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartitionDisjoint(t *testing.T) {
	nums := []int{5, 0, 3, 8, 6}
	x := partitionDisjoint(nums)
	assert.Equal(t, 3, x)
}

func TestPartitionDisjoint2(t *testing.T) {
	nums := []int{1, 1, 1, 0, 6, 12}
	x := partitionDisjoint(nums)
	assert.Equal(t, 4, x)
}
