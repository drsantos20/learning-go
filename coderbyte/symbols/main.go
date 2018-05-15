package main

import (
	"fmt"
	"unicode"
)

func SimpleSymbols(str string) string {

	for i := 0; i < len(str); i++ {

		if IsLetter(string(str[i])) {
			if i == 0 || i == len(str) {
				return "false"
			} else if string(str[i-1]) != "+" || string(str[i+1]) != "+" {
				return "false"
			}
		}
	}

	return "true"
}

func IsLetter(str string) bool {
	for _, char := range str {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(SimpleSymbols("aaaa"))

}

//+d===+a+
//+d+=3=+s+
//aaaa
//2+a+a+
