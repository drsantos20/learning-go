package main

import (
	"fmt"
	"strings"
)

func LetterCapitalize(str string) string {
	upperCaseWord := ""

	for _, word := range strings.Split(str, " ") { //argument
		if len(upperCaseWord) > 0 {
			upperCaseWord += " "
		}
		for i := 0; i < len(word); i++ {
			if i == 0 {
				upperCaseWord += strings.ToUpper(string(word[i]))
				//print(strings.ToUpper(string(word[i])))
			} else {
				upperCaseWord += string(word[i])
				//print(string(word[i]))
			}
		}
	}
	str = upperCaseWord
	return str

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(LetterCapitalize(readline()))

}
