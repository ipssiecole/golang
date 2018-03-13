package main

import "fmt"

func GetNames() (string, string) {
	return "John", "Doe"
}

func ShowNames(firstname, lastname string) {
	fmt.Println(firstname, lastname)
}

type Funky func()

func (f Funky) Dance() {
	f()
}

type Dancer interface {
	Dance()
}

func StartParty(d Dancer) {
	d.Dance()
}

func main() {
	ShowNames(GetNames())

	var f Funky
	f = func() {
		fmt.Println("I'm dancing, I Got Funky")
	}

	StartParty(f)
}
