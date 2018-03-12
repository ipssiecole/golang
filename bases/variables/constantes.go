package main

import (
	"fmt"
)

const language = "Go"

// a == 0
// b == 1
// c == 2
const (
	a = 512 << iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(language, a, b, c, d, e)
}
