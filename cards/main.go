package main

func main() {
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("mycards")
	// cards := newDeckFromFile("mycards")

	cards := newDeck()
	// fmt.Println(cards)
	// fmt.Println(cards.toString())
	// fmt.Println(strings.Join([]string(cards), ","))
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamends"
}
