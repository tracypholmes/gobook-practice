# Chapter 3 Problems

1. How are integers stored on a computer?
    * as bytes

1. We know that (in base 10) the largest 1 digit number is 9 and the largest 2 digit number is 99. Given that in binary the largest 2 digit number is 11 (3), the largest 3 digit number is 111 (7) and the largest 4 digit number is 1111 (15) what's the largest 8 digit number? (hint: 101-1 = 9 and 102-1 = 99)

1. Although overpowered for the task you can use Go as a calculator. Write a program that computes 321325 × 424521 and prints it to the terminal. (Use the * operator for multiplication)

    ```go
    package main

    import "fmt"

    func main() {
      fmt.Println(321325 * 424521)
    }
    ```

1. What is a string? How do you find its length?
    * A string is a sequence of characters of a specific length used to denote text. They are made up of bytes.
    * You find a string's lenth by using `len` preceding the string (in parens). Ex: `len("Dude I'm tired")

1. What's the value of the expression (true && false) || (false && true) || !(false && false)?
    * true