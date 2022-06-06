//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its length and width
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

type Rectangle struct {
	Length int
	Width  int
}

func calculateArea(data Rectangle) int {
	return data.Length * data.Width
}

func calculatePerimeter(data Rectangle) int {
	return (data.Length + data.Width) * 2
}

func main() {
	rectangle := Rectangle{Length: 20, Width: 12}

	area := calculateArea(rectangle)
	fmt.Println("Area is", area)

	perimeter := calculatePerimeter(rectangle)
	fmt.Println("Perimeter is", perimeter)

	rectangle.Length = rectangle.Length * 2
	rectangle.Width = rectangle.Width * 2

	area = calculateArea(rectangle)
	fmt.Println("Area is", area)

	perimeter = calculatePerimeter(rectangle)
	fmt.Println("Perimeter is", perimeter)
}
