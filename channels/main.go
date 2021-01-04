package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
	// "sync"
)

// var wg sync.WaitGroup

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://godoc.org",
		"http://amazon.com",
	}

	// wg.Add(len(links))

	c := make(chan string)

	for _, v := range links {
		go requester(v, c)
	}

	// wg.Wait()
	// for i := 0 ; i < len(links); i++ {
	// 	go requester(<-c , c)
	// }

	for l := range c {
		go func(s string) {
			time.Sleep(4 * time.Second)
			requester(s, c)
		}(l)
		fmt.Println(runtime.NumGoroutine())
	}
}

func requester(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(err, l)
		c <- l
		return
	}
	fmt.Println(l, "is up")
	c <- l
	// wg.Done()
}
