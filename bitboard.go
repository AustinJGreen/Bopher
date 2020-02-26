package main

import (
	"math/bits"
)

func BitboardMarbleCount(bb uint64) int {
	return bits.OnesCount64(bb)
}

func BitboardIsValid(bb uint64) bool {
	return bb&uint64(0xFFE0000000000000) == 0 && BitboardMarbleCount(bb) <= MarbleCount
}

func BitboardIsWon(bb uint64) bool {
	return bb == 0x1F000000000000
}

// Gets whether there is a marble at a hole
func BitboardGet(bb uint64, hole int) bool {
	return bb&(uint64(1)<<hole) != 0
}

func BitboardGetRange(bb uint64, from, to int) uint64 {
	return (bb >> from) & ((uint64(1) << (to - from + 1)) - 1)
}

func BitboardGetUnsafe(bb uint64) uint64 {
	return bb & uint64(0xFFFFFFFFFFFF)
}

func BitboardGetSafe(bb uint64) uint64 {
	return bb & uint64(0x1F000000000000)
}

func BitboardSet(bb *uint64, hole int) {
	*bb |= uint64(1) << hole
}

func BitboardRemove(bb *uint64, hole int) {
	// Safe, removes if hole is taken, unchanged otherwise
	//*bb = ^(^*bb | (uint64(1) << hole))
	*bb &^= uint64(1) << hole
}

func BitboardGetMSB(bb uint64) int {
	if bb == 0 {
		return -1
	}
	return 63 - bits.LeadingZeros64(bb)
}

func BitboardGetLSB(bb uint64) int {
	if bb == 0 {
		return -1
	}
	return bits.TrailingZeros64(bb)
}

func BitboardRotate(bb uint64, fromTeam, toTeam int) uint64 {
	unsafeHoles := BitboardGetUnsafe(bb)
	if fromTeam == toTeam {
		return unsafeHoles
	}

	// Current relative to Yellow
	delta := (fromTeam - toTeam) * 12
	if delta < 0 {
		delta += 48
	}

	unchangedMask := (uint64(1) << (48 - delta)) - 1

	// we only want to rotate the bits (marbles) that are not in
	// the safe area (in the lower 48 bits)
	same := unsafeHoles & unchangedMask

	// Get rotated bits
	rot := bits.RotateLeft64(same, delta)

	// Get bits that were safe (past 48 bits) and re-append back to rotated
	mod := unsafeHoles >> (48 - delta)
	return rot | mod
}

func BitboardPopMSB(bb *uint64) int {
	msb := BitboardGetMSB(*bb)
	*bb &= (uint64(1) << msb) - 1
	return msb
}
