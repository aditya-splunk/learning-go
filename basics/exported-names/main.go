//  Demonstrating exported names

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	// Below line won't work as exported names start with Capital Letter
	// try uncommenting & read the error message
	//fmt.Println(math.pi)
}
