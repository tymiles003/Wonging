package main

import (
	"fmt"
	"github.com/josephyzhou/wonging"
)

func main() {
	fmt.Println("This is the client")

	deck := new(wonging.Deck).Initialize(1)
	deck = deck.Deal()
	deck = deck.Deal()
	deck = deck.DealRandom()
	deck = deck.DealRandom()
	deck = deck.Shuffle()
}
