package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestNotationRollFromString(t *testing.T) {
	result1, err1 := NotationRollFromString("1+6")
	assert.Nil(t, err1)
	assert.Equal(t, *result1, roll{1, 6})

	_, err2 := NotationRollFromString("1 2")
	assert.NotNil(t, err2)

	_, err3 := NotationRollFromString("0+1")
	assert.NotNil(t, err3)

	_, err4 := NotationRollFromString("7+1")
	assert.NotNil(t, err4)

	_, err5 := NotationRollFromString("6+0")
	assert.NotNil(t, err5)

	_, err6 := NotationRollFromString("6+7")
	assert.NotNil(t, err6)

	_, err7 := NotationRollFromString("")
	assert.NotNil(t, err7)

	_, err8 := NotationRollFromString("6 + 6")
	assert.NotNil(t, err8)
}

func TestNotationHoleFromString(t *testing.T) {
	assert.Equal(t, 0, NotationHoleFromString("y1", Yellow))
	assert.Equal(t, 47, NotationHoleFromString("b12", Yellow))
	assert.Equal(t, 0, NotationHoleFromString("r1", Red))
	assert.Equal(t, 47, NotationHoleFromString("y12", Red))
	assert.Equal(t, 0, NotationHoleFromString("g1", Green))
	assert.Equal(t, 47, NotationHoleFromString("r12", Green))
	assert.Equal(t, 0, NotationHoleFromString("b1", Blue))
	assert.Equal(t, 47, NotationHoleFromString("g12", Blue))
	assert.Equal(t, 48, NotationHoleFromString("hy1", Yellow))
	assert.Equal(t, 52, NotationHoleFromString("hy5", Yellow))
	assert.Equal(t, 48, NotationHoleFromString("hr1", Red))
	assert.Equal(t, 52, NotationHoleFromString("hr5", Red))
	assert.Equal(t, 48, NotationHoleFromString("hg1", Green))
	assert.Equal(t, 52, NotationHoleFromString("hg5", Green))
	assert.Equal(t, 48, NotationHoleFromString("hb1", Blue))
	assert.Equal(t, 52, NotationHoleFromString("hb5", Blue))
}

func BenchmarkNotationRollFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		die1 := 1 + rand.Intn(6)
		die2 := 1 + rand.Intn(6)
		rollStr := fmt.Sprintf("%d+%d", die1, die2)
		NotationRollFromString(rollStr)
	}
}
