package main

import (
	"os"
	"fmt"
)

func main() {
	name := "mkz"
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>FROM GO worLD</title>
	</head>
	<body>
	<h1>` + name + os.Args[1] + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
	fmt.Println(os.Args)
}
