package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRoll(t *testing.T) {
	r1 := NewRoll(1, 2)
	assert.NotNil(t, *r1)
	assert.Equal(t, 1, r1.die1)
	assert.Equal(t, 2, r1.die2)
}

func TestRoll_IsDoubles(t *testing.T) {
	r1 := NewRoll(1, 2)
	assert.False(t, r1.IsDoubles())

	r2 := NewRoll(2, 2)
	assert.True(t, r2.IsDoubles())
}
