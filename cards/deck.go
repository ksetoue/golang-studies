package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// creates and returns a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Tree", "Four"}

	for _, suit := range cardSuits { // _ has a variable but I dont care about it
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//receiver --> setups methods on variables we create
// d is the instance of the variable --> similar to this or self
// usually one or two letters --> convention of golang
// deck is a reference to the type 'deck'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// function that returns two types of decks, receiving a deck and a handsize as parameters
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// Option 1 - log error and return call to newDeck()
		// Option2 - log error and quit program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// interesting swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
