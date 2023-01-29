package main

// import (
// 	"fmt"
// )

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)

	// cards := []string{"LMAO",newCard()}
	// cards := newDeck()
	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// greeting :="Aneetta"
	// fmt.Println([]byte(greeting))

	// cards := newDeckFromFile("my")
	// fmt.Println(cards)
	// cards.saveToFile("my_cards")
	// newDeckFromFile("my_cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five DIAMOND"
// } 
