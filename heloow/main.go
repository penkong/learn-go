package main

import "fmt"

func main() {
	// 1 
	colors := map[string]string{
		"red":"5345",
		"blue": "fdsf",
	}

	// 2
	// var colors map[string]string

	// 3
	// colors := make(map[string]string)
	// colors["white"] = "fsdf43t"
	// delete(colors, "white") 
	
	// iterate 

	printMap(colors)
	// fmt.Println(colors)
}


func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}