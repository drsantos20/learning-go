package main

import (
	"fmt"
	"strings"
	"unicode"
)

func LetterChanges(str string) string {

	// code goes here
	// Note: feel free to modify the return type of this function
	s := make([]string, 26)
	s[0] = "A"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "E"
	s[5] = "f"
	s[6] = "g"
	s[7] = "h"
	s[8] = "I"
	s[9] = "j"
	s[10] = "k"
	s[11] = "l"
	s[12] = "m"
	s[13] = "n"
	s[14] = "O"
	s[15] = "p"
	s[16] = "q"
	s[17] = "r"
	s[18] = "s"
	s[19] = "t"
	s[20] = "U"
	s[21] = "v"
	s[22] = "w"
	s[23] = "x"
	s[24] = "y"
	s[25] = "z"
	nextWord := ""

	for _, char := range str {

		if unicode.IsSpace(char) {
			nextWord += " "

		} else if unicode.IsLetter(char) {
			nextWord += getNextValue(string(char), s)
		} else {
			nextWord += string(char)
		}
		str = nextWord
	}

	return str

}

func getNextValue(str string, vSlice []string) string {
	found := ""
	for i := range vSlice {

		if strings.EqualFold(vSlice[i], str) {

			found = vSlice[i+1]
		}
	}
	return found
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(LetterChanges(readline()))

}
