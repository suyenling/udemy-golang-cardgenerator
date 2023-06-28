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

//This function creates and returns a list of playing cards in an array of strings.
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//This function logs the contents of the deck of cards.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//This function creates a 'hand' of cards.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//This function creates a new type of "deck" that is a slice of strings.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//This function is used to create a text file with the name of all the cards in newDeck as strings.
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//This function creates a deck by reading from the hard drive.
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//This logs the error and then entirely quits the program.
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//This part turns a slice of strings into a deck of cards.
	s := strings.Split(string(bs), ",")
	return deck(s)
}

//This function shuffles the cards in the deck.
func (d deck) shuffle() {
	//This part of the code serves as the random generator, using the current time to generate a different int64 number to act as the seed.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		//This part of the code essentially says "take whatever is at 'i' and assign it to 'newPosition,' then take whatever is at 'newPosition' and assign it to 'i.'"
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}