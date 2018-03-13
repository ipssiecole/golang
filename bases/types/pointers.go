package main

import "fmt"

func updater(n *int) {
	*n++
}

func initializer(n *int) {
	*n = 1
}

func main() {

	a := 1
	fmt.Printf("%p\n", &a)

	b := &a
	fmt.Printf("%p\n", b)

	*b = 25

	updater(&a)

	// Affichage de la valeur ?
	fmt.Printf("%d %d\n", a, *b)

	var x, y *int
	x = new(int)
	y = new(int)

	initializer(x)
	initializer(y)
	updater(x)
	updater(y)

	fmt.Println(*x, *y)
}
