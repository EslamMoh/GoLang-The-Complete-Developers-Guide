package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // The 2 versions of card are the same, and both inform go that the card variable is  going to contain some data of type string
	// in the second verions we are relying upon the go compiler to just figure out that variable is supposed to conatin a string card type
	// It does that by reading in the colon equals operator, its better to use :=
	// card := newCard() // We can call a function and insert to a variable
	// cards := deck{"Ace of Diamonds", newCard()} // This is how we declare variable of type deck
	// cards = append(cards, "Six of Spades")      // This will not modify the original slice, instead it will return a new slice with added value

	// cards := newDeck()

	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")

	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
	// for i, card := range cards { // This is how we iterate oer a slice, range is a keyword that we use whenever we want to iterate over every single record inside of a slice
	// 	fmt.Println(i, card)
	// }
}
