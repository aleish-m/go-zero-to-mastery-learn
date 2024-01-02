//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

// * Create a rectangle structure containing a length and width field
type Rectangle struct {
	length int
	width  int
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - The functions must use the rectangle structure as the function parameter

func area(rec Rectangle) int {
	a := rec.length * rec.width
	fmt.Println(a)
	return a
}

func perimeter(rec Rectangle) int {
	p := (rec.length + rec.width) * 2
	fmt.Println(p)
	return p
}

func main() {
	//  - Print the results to the terminal
	rec1 := Rectangle{2, 4}
	area(rec1)
	perimeter(rec1)

	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	rec2 := Rectangle{rec1.length * 2, rec1.width * 2}
	//  - Print the new results to the terminal
	area(rec2)
	perimeter(rec2)

}
