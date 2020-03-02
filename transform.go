package main

type transform struct {
	key        uint64
	removed    int // Int for storing any leading marbles that were removed
	toMove     int
	doublesCnt byte
	prev       *transform
}

func TransformNew() *transform {
	t := new(transform)
	t.key = 0
	t.removed = -1
	t.toMove = 0
	t.doublesCnt = 0
	t.prev = nil
	return t
}

func (t transform) IsValid() bool {
	return t.key >= 0 && t.removed >= -1 && t.removed < 54 && t.toMove >= 0 && t.toMove <= TeamCount && t.doublesCnt <= 2
}
