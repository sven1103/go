package main

import (
	"fmt"
	"math/rand"
)

func Array() {
	var fiveZero [5]int

	fmt.Printf("initial array: %v\n", fiveZero)

	for index, value := range fiveZero {
		fmt.Println(value)
		fiveZero[index] = rand.Int()
	}

	fmt.Printf("array with random int: %v\n", fiveZero)
	fmt.Printf("sliced array: %v\n", fiveZero[0:2])

}
