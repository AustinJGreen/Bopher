package main

import "fmt"

const Yellow = 0
const Red = 1
const Green = 2
const Blue = 3
const TeamCount = 4
const MarbleCount = 5

func NextTeam(team int) int {
	return (team + 1) % TeamCount
}

func main() {
	r := roll{1, 2}
	fmt.Printf("Rolled a %d and %d!\n", r.die1, r.die2)

	cur := Blue
	after := NextTeam(cur)
	fmt.Printf("The team after %d is %d\n", cur, after)
}
