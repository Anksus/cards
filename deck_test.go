package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected deck len of 20 but got %v", len(d))
	}
}
