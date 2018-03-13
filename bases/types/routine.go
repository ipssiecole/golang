package main

import (
	"fmt"
	"time"
)

// Do not communicate by sharing memory;
// instead, share memory by communicating.
func main() {
	chanCounter := make(chan int)

	go func(ch chan int) {
		fmt.Println("A long task")
		counter := <-ch
		counter++
		ch <- counter
	}(chanCounter)

	go func(ch chan int) {
		fmt.Println("A long task")

		counter := <-ch
		counter++
		ch <- counter
	}(chanCounter)

	time.Sleep(time.Second * 3)

	go func(ch chan int) {
		fmt.Println("display")
		for counter := range ch {
			fmt.Println(counter)
		}
		close(ch)

	}(chanCounter)
}
