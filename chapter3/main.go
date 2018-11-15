package main

import (
	"fmt"
)

func main() {
	// fmt.Println("1 + 1 +", 1+1)
	fmt.Println(len("Hello World")) // A space is also considered a character, so the string's length is 11 not 10
	fmt.Println("Hello World"[1])
	fmt.Println("hello " + "World")
}
