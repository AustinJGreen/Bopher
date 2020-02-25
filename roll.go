package main

type roll struct {
	die1 int
	die2 int
}

// Method is just a value receiver, doesn't change roll
// so it doesn't need a *roll
func (r roll) IsDoubles() bool {
	return r.die1 == r.die2
}

func NewRoll(die1, die2 int) *roll {
	r := new(roll)
	r.die1 = die1
	r.die2 = die2
	return r
}
