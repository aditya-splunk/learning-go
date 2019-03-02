package main

import (
	"fmt"
)

// function declaration
func add(x int, y int) int {
	return x + y
}

// shorthand declaration for functions
func add2(x, y int) int {
	return x + y
}

// multiple results can be returned from function
func swap(x, y string) (string, string) {
	return y, x
}

// naked return & named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add2(27, 43))
	fmt.Println(add(10, 33))
	a, b := "hello", "world"
	fmt.Println("Before:", a, b)
	a, b = swap(a, b)
	fmt.Println("After:", a, b)
	fmt.Println(split(27))
}
