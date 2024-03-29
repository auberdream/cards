package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	deckName := "_decktesting"
	os.Remove(deckName)

	d := newDeck()
	d.saveToFile(deckName)

	loadedDeck := newDeckFromFile(deckName)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove(deckName)
}
