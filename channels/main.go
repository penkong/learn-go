package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://godoc.org",
		"http://amazon.com",
	}

	wg.Add(len(links))

	for _, v := range links {
		go requester(v)
	}

	wg.Wait()
}

func requester(l string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(err, l)
		return
	}
	fmt.Println(l, "is up")
	wg.Done()
}
