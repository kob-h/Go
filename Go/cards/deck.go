package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := deck{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := deck{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range(cardSuits){
		for _, val := range(cardValues){
			cards = append(cards, fmt.Sprintf("%s of %s", val, suit))
		}
	} 

	return cards
}
func (d deck) print() {
	for i, elem := range(d) {
		fmt.Println(i, elem)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	res, err := os.ReadFile(filename)
	if err != nil{
		return nil, err
	}

	str := string(res)
	return strings.Split(str, ","), nil
}

func (d deck) shuffle() {
	deckLength := len(d)
	if (deckLength == 1){
		return 
	}
	for i := range d {
		randomIndex := rand.Intn(deckLength)
		//Two ways to swipe:
		//swap(d, i, randomIndex)
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}

func swap(d deck, a int, b int){
	temp := d[a]
	d[a] = d[b]
	d[b] = temp
}