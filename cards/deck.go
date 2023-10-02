package main

import (
	"fmt"
)

type deck []string

// This function creates a new deck.
// This does not need a receiver due to the fact
// That this is creating a new object. By convention, it does not make a lot of sense to do so.

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// This function prints all of the items in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// This function return two decks, one without the elements and one with the removed elements.
// The return annotation, tells us the types of data that will be returned.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
