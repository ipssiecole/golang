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

type Str string

func (s Str) Dance() {
	fmt.Println(s)
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
	StartParty(Str("Je suis une chaîne de caractère et je dance, ça te pose un problème ?"))
}
