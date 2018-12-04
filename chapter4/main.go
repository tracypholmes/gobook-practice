package main

import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	z := "omgimsickashell"
	fmt.Println(z)

	const y string = "oi vey"
	fmt.Println(y)

	// Example program that takes a number from user and doubles it
	fmt.Print("Hey! Enter a number: ") // Prompts user
	var input float64                  // assigns user's input to var and converts to float64
	fmt.Scanf("%f", &input)            // Scanf reads the input & fills it with the user's number

	output := input * 2 // variable output is user's input x2

	fmt.Println(output) // Print out contents of output
}
