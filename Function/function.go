package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

//if function arguments are the same type you can omit type from all but the last.
func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println("28 + 83 =", add(28, 83))
	fmt.Println("8 * 7 =", multiply(8, 7))
}
