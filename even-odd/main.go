package main

import (
	"fmt"
)

func main() {

	integerSlices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, odd := range integerSlices {
		if odd%2 == 0 {
			fmt.Printf("\n %v is even", odd)

		} else {
			fmt.Printf("\n %v is odd", odd)

		}
	}
}
