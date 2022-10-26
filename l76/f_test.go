package l76

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	st := "ABC"
	x := MinWindow(s, st)
	assert.Equal(t, "BANC", x)
}
