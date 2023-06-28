package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	cards.saveToFile("random_cards.txt")
}