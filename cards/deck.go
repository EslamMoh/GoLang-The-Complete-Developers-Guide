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

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // replace unused variable with _
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { //d deck is a receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // deal function should take the existing deck of playing cards, taking a number of cards to deal out, and then we'll creare a new deck of the specified size.

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string { // Convert deck type to string
	// strings.Join is used from the package strings to join array of strings
	return strings.Join([]string(d), ",") // []string(d) => convert deck into slice of strings

}

func (d deck) saveToFile(filename string) error { // choose to return error if package returned an error
	// ioutil.WriteFile is used from a go package called ioutil
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // []byte(d.toString()) => to convert string to slice of bytes // 0666 => is a permission means that anyone can read of write the created file
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // ReadFile will return byteslice and error (if it exists), so we will need to use multireturn, and if no error returned then the value of err will be nil

	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1) // os is Go package, This tells the program to quit entirely
	}
	// We receive byte slice from ReadFile, we need to convert it to deck type
	s := strings.Split(string(bs), ",") // string(bs) => to convert the byte slice to string that contain the values comma sparated, strings.Split() => to convert string into slice

	return deck(s)
}

// Shuffle cards => to implement a func to shuffle slice elements, this is not built in in go
// We can acheive this by implementing a func to iterate over the slice, generate a random number and betweem the 0 and slice length, then swap values based on the generated number
// We will generate random numbers using Match.Intn

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // provid the rand function with current time in unix format to use it as a seed to generate new source object
	r := rand.New(source)                           // we use that source object as the basis for our new random generator

	for i := range d {
		newPosition := r.Intn(len(d) - 1) // len(d) => length of slice d // r: is the seed for the random func

		d[i], d[newPosition] = d[newPosition], d[i] //this will swap the slice elements
	}
}
