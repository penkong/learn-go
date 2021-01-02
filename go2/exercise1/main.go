package main

import (
	"fmt"
	"runtime"
)

// var x int
// var y string
// var z bool

type myType int

var x myType
var y int

func main() {
	// exe1, 2, 3
	// exer1()

	// x = 42
	// y = "James Bond"
	// z = true
	// got := fmt.Sprintf("our there var %v %v %v", x,y,z)
	// fmt.Println(got)

	// fmt.Println(x)
	// fmt.Println(fmt.Sprintf("%T", x))
	// x = 42
	// fmt.Println(x)
	// y = int(x)
	// fmt.Println(y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}

func exer1() {

	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
}
