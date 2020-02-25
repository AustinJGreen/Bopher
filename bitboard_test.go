package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitboardMarbleCount(t *testing.T) {
	var a uint64 = 281474976710655
	assert.Equal(t, 48, BitboardMarbleCount(a))

	var b uint64 = 137455731464
	assert.Equal(t, 5, BitboardMarbleCount(b))

	var c uint64 = 0
	assert.Equal(t, 0, BitboardMarbleCount(c))
}

func TestBitboardIsWon(t *testing.T) {
	var d uint64 = 137455731464
	assert.False(t, BitboardIsWon(d))

	var e uint64 = 8725724278030336
	assert.True(t, BitboardIsWon(e))
}

func TestBitboardGet(t *testing.T) {
	var a uint64 = 5629499534213120
	assert.False(t, BitboardGet(a, 51))
	assert.True(t, BitboardGet(a, 50))
	assert.True(t, BitboardGet(a, 52))
	assert.False(t, BitboardGet(a, 0))
	assert.True(t, BitboardGet(uint64(1), 0))
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

func TestBitboardIsValid(t *testing.T) {
	var a uint64 = 0x20000000000000
	assert.False(t, BitboardIsValid(a))

	var b uint64 = 0x1F000000000000
	assert.True(t, BitboardIsValid(b))

	var c uint64 = 0x1002002008042
	assert.False(t, BitboardIsValid(c))

	var d uint64 = 0
	assert.True(t, BitboardIsValid(d))

	var e uint64 = 0x10000200000034
	assert.True(t, BitboardIsValid(e))
}

func TestBitboardRemove(t *testing.T) {
	var f uint64 = 0x1F000000000000
	BitboardRemove(&f, 47)
	assert.True(t, BitboardIsWon(f))

	BitboardRemove(&f, 48)
	assert.False(t, BitboardIsWon(f))
}

func TestBitboardGetUnsafe(t *testing.T) {
	var aa uint64 = 586472905206924040
	assert.Equal(t, uint64(160528718627592), BitboardGetUnsafe(aa))
}

func TestBitboardGetSafe(t *testing.T) {
	var aa uint64 = 586472905206924040
	assert.Equal(t, uint64(844424930131968), BitboardGetSafe(aa))
}

func TestBitboardGetRange(t *testing.T) {
	var f uint64 = 0x1F000000000000
	assert.Equal(t, uint64(31), BitboardGetRange(f, 48, 52))
	assert.Equal(t, uint64(30), BitboardGetRange(f, 47, 51))
}

func TestBitboardRotate(t *testing.T) {
	var zz uint64 = 0x1D800000000000
	assert.Equal(t, uint64(0x800000000000), BitboardRotate(zz, Yellow, Yellow))

	var v uint64 = 0x2002002002
	assert.Equal(t, uint64(0x2002002002), BitboardRotate(v, Green, Yellow))

	var tt uint64 = 0x1E000000000001
	assert.Equal(t, uint64(0x1000), BitboardRotate(tt, Yellow, Blue))

	var r uint64 = 0x800000000001
	assert.Equal(t, uint64(0x1800), BitboardRotate(r, Yellow, Blue))
	assert.Equal(t, uint64(0x800000000001), BitboardRotate(BitboardRotate(r, Yellow, Blue), Blue, Yellow))
	assert.Equal(t, uint64(0x1800000), BitboardRotate(r, Yellow, Green))
	assert.Equal(t, uint64(0x800000000001), BitboardRotate(BitboardRotate(r, Yellow, Green), Green, Yellow))
	assert.Equal(t, uint64(0x1800000000), BitboardRotate(r, Yellow, Red))
	assert.Equal(t, uint64(0x800000000001), BitboardRotate(BitboardRotate(r, Yellow, Red), Red, Yellow))

	var w uint64 = 0x1F000000000000
	assert.Equal(t, uint64(0), BitboardRotate(w, Yellow, Yellow))
	assert.Equal(t, uint64(0), BitboardRotate(w, Yellow, Red))
	assert.Equal(t, uint64(0), BitboardRotate(w, Yellow, Green))
	assert.Equal(t, uint64(0), BitboardRotate(w, Yellow, Blue))
	assert.Equal(t, uint64(0), BitboardRotate(w, Red, Yellow))
	assert.Equal(t, uint64(0), BitboardRotate(w, Red, Red))
	assert.Equal(t, uint64(0), BitboardRotate(w, Red, Green))
	assert.Equal(t, uint64(0), BitboardRotate(w, Red, Blue))
	assert.Equal(t, uint64(0), BitboardRotate(w, Green, Yellow))
	assert.Equal(t, uint64(0), BitboardRotate(w, Green, Red))
	assert.Equal(t, uint64(0), BitboardRotate(w, Green, Green))
	assert.Equal(t, uint64(0), BitboardRotate(w, Green, Blue))
	assert.Equal(t, uint64(0), BitboardRotate(w, Blue, Yellow))
	assert.Equal(t, uint64(0), BitboardRotate(w, Blue, Red))
	assert.Equal(t, uint64(0), BitboardRotate(w, Blue, Green))
	assert.Equal(t, uint64(0), BitboardRotate(w, Blue, Blue))

	var x uint64 = 0x800000000000
	assert.Equal(t, uint64(0x800), BitboardRotate(x, Yellow, Blue))
	assert.Equal(t, uint64(0x800), BitboardRotate(x, Red, Yellow))
	assert.Equal(t, uint64(0x800), BitboardRotate(x, Green, Red))
	assert.Equal(t, uint64(0x800), BitboardRotate(x, Blue, Green))

	var y uint64 = 0x1000000
	assert.Equal(t, uint64(1), BitboardRotate(y, Yellow, Green))
	assert.Equal(t, uint64(1), BitboardRotate(y, Green, Yellow))
	assert.Equal(t, uint64(1), BitboardRotate(y, Red, Blue))
	assert.Equal(t, uint64(1), BitboardRotate(y, Blue, Red))

	var z uint64 = 0x1000
	assert.Equal(t, uint64(1), BitboardRotate(z, Yellow, Red))
	assert.Equal(t, uint64(1), BitboardRotate(z, Red, Green))
	assert.Equal(t, uint64(1), BitboardRotate(z, Green, Blue))
	assert.Equal(t, uint64(1), BitboardRotate(z, Blue, Yellow))
}

func TestBitboardGetLSB(t *testing.T) {
	var a uint64 = 0x800000001020A
	assert.Equal(t, 1, BitboardGetLSB(a))
	assert.Equal(t, -1, BitboardGetLSB(0))
	assert.Equal(t, 0, BitboardGetLSB(1))

	var b uint64 = 0x8000000000000000
	assert.Equal(t, 63, BitboardGetLSB(b))
}

func TestBitboardGetMSB(t *testing.T) {
	var a uint64 = 0x800000001020A
	assert.Equal(t, 51, BitboardGetMSB(a))
	assert.Equal(t, -1, BitboardGetMSB(0))
	assert.Equal(t, 0, BitboardGetMSB(1))

	var b uint64 = 0x8000000000000001
	assert.Equal(t, 63, BitboardGetMSB(b))
}
