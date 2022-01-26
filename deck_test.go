package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected deck len of 20 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected last card to be Five of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards in deck but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
