package main

import "fmt"

var deckSize = 5

func main() {
	cards := newDeck()

	fmt.Println("Before shuffling...")
	cards.print()
	cards.shuffle()
	fmt.Println("After shuffling!")
	cards.print()

	hand, remainingCards := deal(cards, deckSize)

	fmt.Println("Hand:")
	hand.print()

	fmt.Println("Remaining cards:")
	remainingCards.print()

	fmt.Println("Saving cards...")
	err := cards.saveToFile("data/cards.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Cards saved!")

	fmt.Println("Loading cards...")
	cardsFromFile := newDeckFromFile("data/cards.txt")
	cardsFromFile.print()

	fmt.Println("Cards loaded!")
}
