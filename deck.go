package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
			val += ","
		}
	}
	return val
}

// Custom implementation

// func stringToCards(str string) deck {
// 	cards := deck{}
// 	temp := ""
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == ',' {
// 			cards = append(cards, temp)
// 			temp = ""
// 		} else {
// 			temp += string(str[i])
// 		}
// 	}
// 	if temp != "" {
// 		cards = append(cards, temp)
// 	}
// 	return cards
// }

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	str := string(bs)
	ss := strings.Split(str, ",")
	return deck(ss)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}

}
