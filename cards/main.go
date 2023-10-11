package main

func main() {
	cards := newDeckFromFile("cards")
	cards.shuffle()
	//This invoke selects cards from the deck created above.

	//hand, remainingCards := deal(cards, 4)

	cards.print()
}
