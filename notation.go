package main

import (
	"errors"
	"strconv"
)

var teams = [TeamCount]string{"yellow", "red", "green", "blue"}

func NotationRollFromString(rollStr string) (*roll, error) {
	if len(rollStr) != 3 {
		return nil, errors.New("roll should be in format die1+die2")
	} else if rollStr[1] != '+' {
		return nil, errors.New("bad roll string")
	}
	num1 := int(rollStr[0] - 48)
	num2 := int(rollStr[2] - 48)
	roll := RollNew(num1, num2)
	if num1 < 1 || num1 > 6 || num2 < 1 || num2 > 6 {
		return nil, errors.New("roll out of range")
	}

	return roll, nil
}

func NotationHoleFromString(holeStr string, teamPerspective int) int {
	if holeStr[0] == 'h' {
		if len(holeStr) < 3 {
			return -1
		}

		hole, err := strconv.Atoi(holeStr[2:])
		if err != nil {
			return -1
		}

		return 47 + hole
	}

	for t := 0; t < TeamCount; t++ {
		if holeStr[0] == teams[t][0] {
			value, err := strconv.Atoi(holeStr[1:])
			if err != nil {
				return -1
			}

			return BitboardGetMSB(BitboardRotate(uint64(1<<(value-1)), t, teamPerspective))
		}
	}

	return -1
}
