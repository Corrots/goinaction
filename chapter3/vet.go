package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	fmt.Println(<-ch)

	//fmt.Printf("Hahaha", 3.14)
}
