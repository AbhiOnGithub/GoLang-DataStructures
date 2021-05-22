//main package has examples shown
package main

// importing fmt package
import (
	"fmt"
)

//gets the powerToNumber of integer num and returns tuple of square of num and cube of num
func powerToNumber(num int) (int, int) {
	return num * num, num * num * num
}

//gets the powerToNumberByNamedTuple of integer num and returns tuple of square of num and cube of num
func powerToNumberByNamedTuple(num int) (square int, cube int) {
	square = num * num
	cube = num * num * num
	return square, cube
}

// main method
func main() {
	var square int
	var cube int
	square, cube = powerToNumber(3)
	fmt.Println("Square ", square, "Cube", cube)

	square, cube = powerToNumberByNamedTuple(5)
	fmt.Println("Square ", square, "Cube", cube)
}
