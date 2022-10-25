package l934

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestBridge(t *testing.T) {
	nums := [][]int{{0,1},{1,0}}
	x := shortestBridge(nums)
	assert.Equal(t, 1, x)
}
