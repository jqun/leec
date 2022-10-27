package l30

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	x := findSubstring(s, words)
	assert.Equal(t, []int{0, 9}, x)
}

func TestFindSubstring2(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word","good","best","word"}
	x := findSubstring(s, words)
	var y []int
	assert.Equal(t, y, x)
}

func TestFindSubstring3(t *testing.T) {
	s := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	words := []string{"fooo","barr","wing","ding","wing"}
	x := findSubstring(s, words)
	assert.Equal(t, []int{13}, x)
}

func TestFindSubstring4(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	x := findSubstring(s, words)
	assert.Equal(t, []int{8}, x)
}
