package main

type position struct {

	current *transform
	teams   [TeamCount]uint64
}

func PositionNew() *position {
	p := new(position)
	p.current = nil

	for i := 0; i <= TeamCount; i++ {
		p.teams[i] = 0
	}

	return p
}

func (p *position) SetByName(sq string, team int) {
	p.SetByHole(NotationHoleFromString(sq, team), team)
}

func (p *position) SetByHole(hole, team int) {
	BitboardSet(&p.teams[team], hole)
}

func (p *position) RemoveLeading(team int) int {
	// assert(magic_isValid(teams[team]))
	unsafe := BitboardGetUnsafe(p.teams[team])
	if unsafe == 0 {
		return -1
	}

	leading := BitboardGetMSB(unsafe)
	BitboardRemove(&p.teams[team], leading)

	return leading
}

func (p *position) Reset() {
	for i := 0; i < TeamCount; i++ {
		p.teams[i] = 0
	}

	p.current = TransformNew()
}

func (p *position) Next(r roll) {
	next := TransformNew()
	next.doublesCnt = p.current.doublesCnt
	if r.IsDoubles() && p.current.doublesCnt + 1 < 3 {
		next.toMove = p.current.toMove
		next.doublesCnt++
	} else {
		if p.current.doublesCnt + 1 >= 3 {
			r := p.RemoveLeading(p.current.toMove)
			next.removed = r
		}

		next.toMove = NextTeam(p.current.toMove)
		next.doublesCnt = 0
	}
}