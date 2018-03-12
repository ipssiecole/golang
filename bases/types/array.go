package main

import "fmt"

func main() {
	// les arrays (taille fixe)

	arr := [5]string{
		"John",
		"Franck",
		"Jenny",
		"Harry",
	}

	arr2 := [...]int{8, 10, 12, 89}

	fmt.Println(arr, arr2)
}
