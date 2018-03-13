package main

import "fmt"

type user struct {
	firstname string
	lastname  string
	age       uint8
}

func (v user) sayHello() string {
	return "Hello " + v.firstname
}

// Créer un type Product struct (name, description, price)
// une fonction display() qui retourne le nom et la description

type product struct {
	name        string
	description string
	price       float64
}

func (p product) display() string {
	return p.name + " " + p.description
}

func main() {
	u := user{
		firstname: "John",
		lastname:  "Doe",
		age:       30,
	}
	u.age = 100
	fmt.Println(u.sayHello())
	fmt.Println(u.firstname, u.lastname, u.age)

	p := product{
		name:        "Iphone",
		description: "no comment",
		price:       10000.0,
	}
	fmt.Println(p.display())
}
