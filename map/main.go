package main

import (
	"fmt"
)

func main() {

	// var colour map[string]string

	// It will create empty map and let the value to be assigned later.
	// colour := make(map[string]string)
	colour := map[string]string{
		"red":   "#12312",
		"green": "#46788",
	}

	colour["red"] = "213311222333"
	// a := colour["red"]

	// fmt.Println(a)
	mapPrint(colour)
}

func mapPrint(c map[string]string) {

	for colour, _ := range c {
		if colour == "red" {
			fmt.Println("there is red colour in map")
		}
	}
}
