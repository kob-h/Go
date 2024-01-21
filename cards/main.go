package main

import (
	"fmt"
)

func main(){
	
	de := newDeck()
	fmt.Println("Before shuffling")
	de.print()
	de.shuffle()
	fmt.Println("after shuffling")
	de.print()
}

func newCard() string {
	return "Ace of Spades"
}