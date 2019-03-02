package main

// variables can be at a package level or function level

import (
	"fmt"
)

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
