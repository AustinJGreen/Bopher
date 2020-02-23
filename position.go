package main

type position struct {
	current *transform
	teams   [TeamCount]uint64
}

func NewPosition() *position {
	p := new(position)
	p.current = nil

	for i := 0; i <= TeamCount; i++ {
		p.teams[i] = 0
	}

	return p
}

func (p *position) SetByName(sq string, team int) {

}

func (p *position) SetByHole(team, hole int) {

}
