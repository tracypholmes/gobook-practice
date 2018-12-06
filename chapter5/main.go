package main

import "fmt"

/*
// option 1
func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
*/

// Option 2

func main() {
	for i := 1; i <= 13; i++ {
		if i%2 == 0 { // Is the remainder of i ÷ 2 equal to 0? No: jump to the else block
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	t := 9 // wasn't working because I didn't declare variable first!
	switch t {
	case 0:
		fmt.Println("Zed")
	case 1:
		fmt.Println("One ha ha")
	case 2:
		fmt.Println("TWO HA HA")
	default:
		fmt.Println("Unknown Number - retreat!")
	}
}
