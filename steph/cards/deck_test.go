package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected deck length of 20, but got %d\n",len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("execpted but not %v\n", d[0])
	}

	if d[len(d) - 1] != "Clubs of Four" {
		t.Errorf("bad")
	}

}


func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("error in loaded file")
	}
	os.Remove("_decktesting")
}