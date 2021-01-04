package main





func main(){
	// var card string = "Ace of space"

	// card := newCard()

	// reassign
	// card = "Five of Diamonds"

	// fmt.Println(card)

	// slice
	// cards := deck{newCard(), newCard(), "sth"}

	// add to slice 
	// cards = append(cards, "Six of sth")

	// loop over slice 
	// for i, card := range cards {
	// 	fmt.Println(i,card)
	// }
	
		// type conversion
		// greeting := "H i there!"
		// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	
	// cards := newDeck()
	// cards.saveToFile("my_cart")

	cards := newDeckFromFile("my_cart")
	cards.shuffle()
	cards.print()
}



// func newCard() string {
// 	return "Five of Diamonds"
// }