package main

import (
	"fmt"
)

func main() {
	// success buffer
	c := make(chan int, 2)
	// go func() {
	// 	c <- 42
	// 	}()
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
}
