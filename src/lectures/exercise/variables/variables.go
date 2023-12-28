//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import (
	"fmt"
)

func main() {
	var favColor = "teal"
	fmt.Println("My favorite color is", favColor)

	yearBorn, age := 2000, 30
	fmt.Println("If I was born in", yearBorn, ". Would I be", age, "years old?")

	var (
		firstInitial = "A"
		lastInitial  = "M"
	)
	fmt.Println("My initials are", firstInitial, lastInitial)

	var ageInDays int

	ageInDays = age * 365

	fmt.Println(age, "years is", ageInDays, "days.") // This concatenates when printed with spaces between values and a new line.
}
