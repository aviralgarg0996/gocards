package main
import "fmt"

type deck []string

func newDeck() deck {
    cardSuits :=[]string{"Spades","Diamonds","Hearts","Clubs"}
    cardValues:=[]string{"Ace","Two","Three","Four"}
    cards := deck{}

    for _ ,suit :=range cardSuits{
        for _ ,value :=range cardValues{
            cards=append(cards,value+" of "+suit)
        }
    }
    return cards
}

func (d deck) print(){
    for i, card := range d{
        fmt.Println(i,card)
    }
}

func deal(d deck ,handSize int) (deck,deck){
    return d[:handSize] , d[handSize:]

}
