package main

type roll struct {
	die1 int
	die2 int
}

func (r roll) isDoubles() bool {
	return r.die1 == r.die2
}
