package l567

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	b := checkInclusion(s1, s2)
	assert.Equal(t, true, b)
}

func TestCheckInclusion2(t *testing.T) {
	s1 := "ab"
	s2 := "eidboaoo"
	b := checkInclusion(s1, s2)
	assert.Equal(t, false, b)
}
