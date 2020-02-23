package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitboardMarbleCount(t *testing.T) {
	var a uint64 = 281474976710655
	assert.Equal(t, BitboardMarbleCount(a), 48)

	var b uint64 = 137455731464
	assert.Equal(t, BitboardMarbleCount(b), 5)

	var c uint64 = 0
	assert.Equal(t, BitboardMarbleCount(c), 0)
}

func TestBitboardIsWon(t *testing.T) {
	var d uint64 = 137455731464
	assert.False(t, BitboardIsWon(d))

	var e uint64 = 8725724278030336
	assert.True(t, BitboardIsWon(e))
}

func TestBitboardSet(t *testing.T) {
	var f uint64 = 0
	BitboardSet(&f, 48)
	BitboardSet(&f, 49)
	BitboardSet(&f, 50)
	BitboardSet(&f, 51)
	BitboardSet(&f, 52)
	assert.True(t, BitboardIsWon(f))
}

func TestBitboardRemove(t *testing.T) {
	var f uint64 = 8725724278030336
	BitboardRemove(&f, 47)
	assert.True(t, BitboardIsWon(f))

	BitboardRemove(&f, 48)
	assert.False(t, BitboardIsWon(f))
}
