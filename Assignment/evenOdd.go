package main

import (
	"fmt"
)

func genrateNumber(number int) []int {

	arr := []int{}
	for i := 1; i <= number; i++ {
		arr = append(arr, i)
	}

	return arr
}

func TestEvenOrOdd(number []int) {

	for i := 0; i <= len(number)-1; i++ {
		if number[i]%2 == 0 {
			// fmt.Println(number[i], "is even")
			fmt.Println(fmt.Sprintf("%d is even", number[i]))
		} else {
			fmt.Println(number[i], "is odd")
		}
	}
}
