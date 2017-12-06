package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// If function arguments are the same type you can omit type from all but the last.
func multiply(x, y int) int {
	return x * y
}

// Function can return multiple results.
func swap(x, y string) (string, string) {
	return y, x
}

// Function can have "naked" return.
// x and y are treated as variables defined at the top of function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("28 + 83 =", add(28, 83))

	fmt.Println("8 * 7 =", multiply(8, 7))

	a, b := swap("Hello", "world")
	fmt.Println("Hello World after swap: ", a, b)

	fmt.Println(split(17))
}
