package main

import "fmt"

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// println()
	// remainingCards.print()
	// println(hand.toString())
	// cards.saveToFile("cardNames")
	cards := newDeckFromFile("cardNames")
	cards.print()
	cards.shuffle()
	fmt.Println("---------")
	cards.print()
}
