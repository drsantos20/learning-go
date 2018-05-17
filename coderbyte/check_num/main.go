package main

import "fmt"

func CheckNums(num1 int, num2 int) int {

	// code goes here
	// Note: feel free to modify the return type of this function
	if num2 > num1 {
		return 1
	} else if num2 == num1 {
		return -1
	} else {
		return 0
	}

	return num1

}

func TypeReturn(num int) string {
	if num == 1 {
		return "true"
	} else if num == 0 {
		return "false"
	} else {
		return "-1"
	}
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(TypeReturn(CheckNums(readline())))

}
