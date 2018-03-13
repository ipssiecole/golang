package main

import "fmt"

func main() {

	m := map[string]string{
		"name":    "John",
		"country": "UK",
	}

	if value, ok := m["name"]; ok == true {
		fmt.Println(value, m["country"])
	}

	var a interface{}
	a = 5
	b := 10

	fmt.Println(a.(int) + b)

	var x interface{}
	x = "Une chaîne"

	fmt.Println(x.(string) + " Une autre chaîne")

	m2 := map[string]interface{}{
		"prenom": "John",
		"nom":    "Smith",
		"age":    35,
		"size":   2.66,
	}

	fmt.Println(m2["prenom"])
}
