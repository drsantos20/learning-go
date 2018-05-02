package main

import "fmt"

func SimpleAdding(num int) int {

	// code goes here
	// Note: feel free to modify the return type of this function
	sum := 0
	for i := 0; i <= num; i++ {

		sum += i

	}

	num = sum
	return num

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(SimpleAdding(readline()))

}
