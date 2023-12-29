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

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func sayHello(name string) {
	fmt.Println("Hello", name, "!")
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func saySomething() string {
	return "Figuring this language out. The new terminal makes this harder... but I will get comfortable with this soon!"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func addThree(x, y, z int) int {
	return x + y + z
}

// * Write a function that returns any number
func getOneNum() int {
	return 10
}

// * Write a function that returns any two numbers
func returnTwo(x, y int) (int, int) {
	return x, y
}

func main() {
	sayHello("Bob")
	fmt.Println(saySomething())
	println(addThree(3, 10, 20)) // Sprintf, errorf

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	a, b := returnTwo(5, 10)
	fmt.Println(addThree(getOneNum(), a, b))
}
