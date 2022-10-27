package l3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	x := lengthOfLongestSubstring(s)
	assert.Equal(t, 3, x)
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	s := "bbbbb"
	x := lengthOfLongestSubstring(s)
	assert.Equal(t, 1, x)
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	s := "pwwkew"
	x := lengthOfLongestSubstring(s)
	assert.Equal(t, 3, x)
}
