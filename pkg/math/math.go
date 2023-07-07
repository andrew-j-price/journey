package math

import (
	"fmt"
	"os"

	helper "github.com/andrew-j-price/journey/pkg/helpers"
)

// Add sums two integers
func add(x, y int) int {
	return x + y
}

// Subtract subtracts two integers
func subtract(x, y int) int {
	return x - y
}

func Main(args []string) {
	int1 := helper.StringToInt(args[1], true)
	int2 := helper.StringToInt(args[2], true)
	var result int
	if args[0] == "add" {
		result = add(int1, int2)
	}
	if args[0] == "subtract" {
		result = subtract(int1, int2)
	}
	fmt.Printf("%v result: %v \n", args[0], result)
	os.Exit(0)
}
