package main

import "fmt"

func test(s []string) {
	s[0] = "mangue"
}

func main() {
	s := []string{"kiwi", "pomme", "fraise", "poire"}
	// s[low:high]

	fmt.Println(s[1:2])

	// make(type, size, capacity)

	myslice := make([]int, 0, 2)
	myslice = append(myslice, 52)
	myslice = append(myslice, 48)
	myslice = append(myslice, 14)

	fmt.Println(len(myslice), cap(myslice))
	fmt.Println(myslice)

	str := "Hello World"
	fmt.Println(str[0:5])
}
