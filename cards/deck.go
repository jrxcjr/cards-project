package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

// This function creates a new deck object.
// This does not need a receiver because it is creating a new object
// return @deck cards
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

// This function prints all of the items in a deck to the console
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// This function return two decks, one without the elements and one with the removed elements.
// The return annotation, tells us the types of data that will be returned.
// return @hand deck
// return @remainingCards deck
func deal(d deck, handSize int) (hand deck, remainingCards deck) {

	hand = d[:handSize]
	remainingCards = d[handSize:]

	return hand, remainingCards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) toFileByte() []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, d.toFileByte(), 0666)
}

func newDeckFromFile(filename string) deck {

	bs, err := os.ReadFile(filename)

	if err != nil {
		//option 1- when that happens, we can printout a new deck and log the error.
		//option 2 -Log the error and entirely quit the runtime execution.
		fmt.Println("Error:", err)
		//This kills the execution of a program
		os.Exit(1)
	}

	//this got the byte slice and turned it back to a []strings
	readByteSlice := strings.Split(string(bs), ",")

	//the []string was then converted to a deck type
	return deck(readByteSlice)
}
