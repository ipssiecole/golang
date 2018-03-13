package main

import "fmt"

type Str string

func (s *Str) Concat(value string) {
	*s += Str(value)
}

func main() {
	var myString Str
	myString = "Hello"
	myString.Concat("Gopher")
	myString.Concat(" Super")

	name := "Yoman"
	myName := Str(name)
	myName.Concat("!!!")

	fmt.Println(myString, myName)
}
