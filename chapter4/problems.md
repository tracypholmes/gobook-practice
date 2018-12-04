# Chapter 4 Problems

1. What are two ways to create a new variable?  

    * `var x = "woot"
    * x := "woot


2. What is the value of x after running: `x := 5; x += 1?`

3. What is scope and how do you determine the scope of a variable in Go?

4. What is the difference between var and const?

    * You can reassign a `var` but a `const` can only be assigned once.

5. Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)

```go
package main

import "fmt"

func main() {
  fmt.Print("Enter a temp: ")
  var input float64
  fmt.Scanf("%f", &input)
     
  output := (input - 32) * 5/9
     
  fmt.Println(output)
}
```

6. Write another program that converts from feet into meters. (1 ft = 0.3048 m)