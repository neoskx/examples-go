package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("This is the time: ", time.Now())

	fmt.Println("Random Number: ", rand.Intn(10.00))

	a := add(4, 5)

	fmt.Println(a)

	b, c := swap("string b", "string c")

	fmt.Println(b, c)
}
