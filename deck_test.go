package main

import (
	"os"
	"testing"
)

var deckTestFile = "data/_deck_test.txt"

func TestNewDeck(t *testing.T) {
	// arrange
	expectedDeckCount := 16

	// act
	cards := newDeck()

	// assert
	if len(cards) != expectedDeckCount {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckCount, len(cards))
	}

	expectFirstCard := "Ace of Spades"
	if cards[0] != expectFirstCard {
		t.Errorf("Expected first card of %v, but got %v", expectFirstCard, cards[0])
	}

	expectLastCard := "Four of Clubs"
	if cards[len(cards)-1] != expectLastCard {
		t.Errorf("Expected last card of %v, but got %v", expectLastCard, cards[len(cards)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// arrange
	os.Remove(deckTestFile)
	deck := newDeck()
	expectedDeckCount := 16

	// act
	deck.saveToFile(deckTestFile)

	// assert
	if _, err := os.Stat(deckTestFile); os.IsNotExist(err) {
		t.Errorf("Expected file %v to exist", deckTestFile)
	}

	// act
	loadedDeck := newDeckFromFile(deckTestFile)

	// assert
	if len(loadedDeck) != expectedDeckCount {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckCount, len(loadedDeck))
	}

	os.Remove(deckTestFile)
}
