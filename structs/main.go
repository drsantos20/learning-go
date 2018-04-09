package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	daniel := person{firstName: "Daniel", lastName: "Santos"}
	fmt.Println(daniel)
}
