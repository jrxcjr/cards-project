package main

func main() {
	cards := newDeck()

	//This invoke selects cards from the deck created above.

	hand, remainingCards := deal(cards, 4)

	hand.print()
	remainingCards.print()
}
