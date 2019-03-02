package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("My Favourite Number (This is constant): ", rand.Intn(100))
	seed := rand.NewSource(time.Now().Unix())
	r1 := rand.New(seed)
	fmt.Println("My Favourite Number (This always changes): ", r1.Intn(100))
}
