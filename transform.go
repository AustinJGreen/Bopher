package main

type transform struct {
	key        uint64
	removed    int // Int for storing any leading marbles that were removed
	toMove     int
	doublesCnt byte
	prev       *transform
}

func (t transform) isValid() bool {
	return t.key >= 0 && t.removed >= -1 && t.removed < 54 && t.toMove >= 0 && t.toMove <= TeamCount && t.doublesCnt <= 2
}
