package main

func main() {
	cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// println()
	// remainingCards.print()
	// println(hand.toString())
	cards.saveToFile("cardNames")
}
