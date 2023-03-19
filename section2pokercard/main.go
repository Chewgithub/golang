package main

import "fmt"

func main() {
	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
