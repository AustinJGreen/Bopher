package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestNotationRollFromString(t *testing.T) {
	result1, err1 := NotationRollFromString("1+6")
	assert.Equal(t, err1, nil)
	assert.Equal(t, *result1, roll{1, 6})

	_, err2 := NotationRollFromString("1 2")
	assert.NotEqual(t, err2, nil)

	_, err3 := NotationRollFromString("0+1")
	assert.NotEqual(t, err3, nil)

	_, err4 := NotationRollFromString("7+1")
	assert.NotEqual(t, err4, nil)

	_, err5 := NotationRollFromString("6+0")
	assert.NotEqual(t, err5, nil)

	_, err6 := NotationRollFromString("6+7")
	assert.NotEqual(t, err6, nil)

	_, err7 := NotationRollFromString("")
	assert.NotEqual(t, err7, nil)

	_, err8 := NotationRollFromString("6 + 6")
	assert.NotEqual(t, err8, nil)
}

func TestNotationHoleFromString(t *testing.T) {

}

func BenchmarkNotationRollFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		die1 := 1 + rand.Intn(6)
		die2 := 1 + rand.Intn(6)
		rollStr := fmt.Sprintf("%d+%d", die1, die2)
		NotationRollFromString(rollStr)
	}
}
