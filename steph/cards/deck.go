package main

import (
	"time"
	"math/rand"
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

// what a deck is and how it behave like class
// create a new type
type deck []string


func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}

	return cards
}


func (d deck) print() {
	for i,card := range d {
		fmt.Println(i,card)
	}
}


func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}


func (d deck) toString() string {
	// type conversion
	return strings.Join([]string(d), ",")
	
}


func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(
		filename, 
		[]byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bslice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:" , err)
		os.Exit(1)
	}
	// type conversion
	s := strings.Split(string(bslice), ",")

	return deck(s)
}


func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}