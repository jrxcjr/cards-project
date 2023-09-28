package main

import "fmt"

// this is a declaration of a new type.
// create a new type of 'deck'
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
