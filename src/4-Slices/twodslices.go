package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {
	var rows int
	var cols int

	rows = 7
	cols = 9

	var twodslices = make([][]int, rows)
	var i int

	for i = range twodslices {
		twodslices[i] = make([]int, cols)
	}

	fmt.Println(twodslices) // all 0s

	i = 0
	for i = range twodslices {
		var j = 0
		for j = range twodslices[i] {
			twodslices[i][j] = i * j
		}
	}

	fmt.Println(twodslices) // updated values
}
