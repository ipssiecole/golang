package main

import "log"

func MyFunc() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	panic("Panic à bord")
}

func main() {
	MyFunc()
}
