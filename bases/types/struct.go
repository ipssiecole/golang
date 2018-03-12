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

func main() {
	u := user{
		firstname: "John",
		lastname:  "Doe",
		age:       30,
	}

	u.age = 100

	fmt.Println(u.sayHello())
	fmt.Println(u.firstname, u.lastname, u.age)
}
