package l438

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	x := findAnagrams(s, p)
	assert.Equal(t, []int{0, 6}, x)
}

func TestFindAnagrams2(t *testing.T) {
	s := "abab"
	p := "ab"
	x := findAnagrams(s, p)
	assert.Equal(t, []int{0, 1, 2}, x)
}
