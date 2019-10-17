package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 3)

	remainingCards.print()
	hand.print()
}
