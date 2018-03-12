package main

import "fmt"

var (
	ecole   = "ipssi"
	adresse = "252 rue Claude Tilier"
)

func main() {
	// Déclaration de variables avec un typage statique.
	var name string
	name = "Gopher"

	var x, y int
	x = 10

	fmt.Println(name, x, y)

	// Déclaration de variable avec un typage dynamique.
	a := 60 // équivalent à var a int = 60
	a = 20

	w, z := 5, false

	// Le Swap
	u, v := 10, 35
	v, u = u, v

	fmt.Println(a, w, z, v)

	version := string("ma valeur")

	fmt.Println(version)
}
