package main

import "fmt"

func main() {
	fmt.Println("hello, everyone ")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	boo()
}

func foo() {
	fmt.Println("i am in foo")
}

func boo() {
	fmt.Println("exited")
}
