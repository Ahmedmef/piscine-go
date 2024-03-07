package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	x := 0
	for z := 0; z < 4; z++ {
		fmt.Printf("Player %v: %v, %v, %v\n", z+1, deck[x], deck[x+1], deck[x+2])
		x = x + 3
	}
}
