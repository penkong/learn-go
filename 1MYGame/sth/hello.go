package main

import (
	"fmt"
	"github.com/penkong/gomod"
	"github.com/penkong/gomod/sth"
)

func main() {
	fmt.Println("hellow")
	s := gomod.Hello("mk")
	sth.WhyNot()
	fmt.Println(s)
}