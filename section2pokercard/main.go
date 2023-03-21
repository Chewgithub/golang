package main

func main() {
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
