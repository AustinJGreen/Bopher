package main

import "fmt"

const TeamCount = 4

func main() {
	r := roll{1, 2}
	fmt.Printf("Rolled a %d and %d!", r.die1, r.die2)
}
