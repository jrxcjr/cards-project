package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

// This function shuffles the deck by randomizing the
func (d deck) shuffle() {

	//The int64 seed that will be used to generate the randomness, can be further randomized by passing to the seed function, a amount of nanoseconds elapsed
	//since the epoch, so, whenever we initialize the program, we get a new time (as more time has passed since the epoch) and a new seed number, thus improving randomization.
	// to get truly random numbers, you need to create a new source of random numbers, that is created by calling this function below.
	source := rand.NewSource(time.Now().UnixNano())

	//A new source is then uploaded to a new instance of the Rand type, that will start generating the random numbers
	r := rand.New(source)

	for i := range d {
		//len (d) - 1 returns the length of a slice.
		//By pointing to the newly created Rand type, you can get truly random numbers
		newPosition := r.Intn(len(d) - 1)

		//This piece of code, lets you swap the item position on a slice, by telling the item, what is his new index.
		d[i], d[newPosition] = d[newPosition], d[i]
	}
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
