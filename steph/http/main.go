package main

import (
	"io"
	"fmt"
	"net/http"
	"os"
)

type logWriter struct {}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	// bs := make([]byte , 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	// io.Copy(os.Stdout, resp.Body)

}


func (l logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println(len(bs))
	return len(bs), nil
}