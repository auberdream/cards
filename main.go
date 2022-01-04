package main

func main() {
	cards := newDeck()
	hand, deck := deal(cards, 5)

	hand.print()
	deck.print()
}
