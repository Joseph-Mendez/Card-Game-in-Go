// Joseph Mendez 2018
// Go Card Game
package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}