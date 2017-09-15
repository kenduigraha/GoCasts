package main

import "fmt"

// var card string = "Card" -> it can
// card := "Card" -> it can't
// var card int -> bisa declare di global

func main() {
	// var card string = "King"
	// card := "Card"
	// card = "Kartu"

	// test := 1 // it can string
	// fmt.Println(test)

	// var card int
	// card = 1

	// assign function to variable
	var card string = newCard()

	fmt.Println(card)
}
