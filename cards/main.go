package main

func main() {
	cards := newDeck()

	//This invoke selects cards from the deck created above.
	deal(cards, 4)
	cards.print()
}
