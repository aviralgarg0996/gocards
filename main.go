package main
import "fmt"

func main(){
cards := newDeck()
cards1 :=newDeckFromFile("mycards")
cards.shuffle()
cards1.print()
// hand,remainingCards :=deal(cards,5)
 cards.toString()
 cards.saveToFile("mycards")
// hand.print()
// remainingCards.print()

// greeting :="Hi There"
fmt.Println(cards.toString())

}

func newCard() string{
	return "hello"
}