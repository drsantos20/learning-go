package main

import (
	"fmt"
	"strings"
)

func LongestWord(sen string) string {

	// code goes here
	// Note: feel free to modify the return type of this function
	result := strings.Split(sen, " ")
	size := len(result[0])

	for i := range result {

		if size <= len(result[i]) {
			sen = result[i]
		}
	}

	return sen

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(LongestWord(readline()))

}
