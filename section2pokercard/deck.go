package main

import "fmt"

// Create a new type of deck, which is a slice of strings

type deck []string

// Receiver function, the receiver sets up method on variables that we create
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
