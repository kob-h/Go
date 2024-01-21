package main

import "testing"

func TestNewDeck(t *testing.T){
	d := newDeck()

	deckLength := len(d)

	if (deckLength != 52){
		t.Errorf("Expected deck size of 52, but got %d", deckLength)
	}
}