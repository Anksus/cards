package main

import (
	"fmt"
	"io/ioutil"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamond", "Heart", "Clubs"}
	cardNum := []string{"One", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardNum {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	val := ""
	for i, v := range d {
		val += v
		if i != (len(d) - 1) {
			val += " ,"
		}
	}
	return val
}

func (d deck) saveToFile(fileName string) {
	ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
