package main

import (
	"fmt"
)

func main() {
	// success buffer
	// c := make(chan int, 2)
	// go func() {
	// 	c <- 42
	// 	}()
	// c <- 42
	// c <- 43
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// one way channel
	// channels only send values on it
	// c := make(chan <- int , 2)
	// c <- 43
	// c <- 434

	// channel only receive
	// c := make(<-chan  int , 2)

	// c := make(chan int)
	// go fooSend(c)
	// barReceive(c)
	// for v := range c {
	// 	fmt.Println(v)
	// }


	// select pattern
	// eve := make(chan int)
	// odd := make(chan int)
	// quit := make(chan int)
	// go sendSelect(eve, odd, quit)
	// receive
	// receiveSelect(eve, odd, quit)

	// Fan it pattern

	fmt.Println("about exit")
}

func receiveSelect(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("eve", v)
		case v := <-o:
			fmt.Println("odd", v)
		case v := <-q:
			fmt.Println("quit", v)
			return
		}
	}
}

func sendSelect(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
	close(q)
}

func fooSend(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func barReceive(c <-chan int) {
	fmt.Println(<-c)
}
