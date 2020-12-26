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

	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}



// func newCard() string {
// 	return "Five of Diamonds"
// }