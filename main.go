package main

func main() {
	cards := newDeckFromFile("cards")

	cards.print()

	cards.shuffle()

	cards.saveToFile("cards")
}
