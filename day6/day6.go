package main

import (
	"fmt"
)

// skipping day5 cause i didnt understood it. I understand this one but i know
// imma be doing it the slow way.
// I dont understand why this is after day5. This is so much easier to get the answer for.
// maybe they want me to search for a matematical solution that is faster?
// i dont expect them to be dumb enough for that so idk.
func main() {
	time := 40929790
	distance := 215106415051100
	var winings []int
	for i := 1; i < time; i++ {
		if i*(time-i) > distance {
			winings = append(winings, i)
		}
	}
	fmt.Println(len(winings))
}
