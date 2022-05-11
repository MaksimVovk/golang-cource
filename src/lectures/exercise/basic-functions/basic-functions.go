//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greetingPerson(person string) {
	fmt.Println("Hello,", person, "!")
}

func displayMessage() string {
	return "hi there!"
}

func add(a, b, c int) int {
	return a + b + c
}

func returnNumber() int {
	return 2
}

func returnTwoNumber() (int, int) {
	return 2, 2
}

func main() {
	greetingPerson("Max")
	fmt.Println(displayMessage())

	sum := add(1, 2, 3)
	fmt.Println("sum =", sum)

	num := returnNumber()

	fmt.Println("num =", num)

	firstNum, secondNum := returnTwoNumber()
	fmt.Println("firstNum =", firstNum, "secondNum =", secondNum)

	result := add(returnNumber(), firstNum, secondNum)
	fmt.Println("result =", result)
}
