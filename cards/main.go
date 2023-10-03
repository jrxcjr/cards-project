package main

func main() {
	cards := newDeck()

	//This invoke selects cards from the deck created above.

	//hand, remainingCards := deal(cards, 4)

	cards.saveToFile("my_cards")
}
