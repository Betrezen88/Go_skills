package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

//if function arguments are the same type you can omit type from all but the last.
func multiply(x, y int) int {
	return x * y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Hello", "world")
	fmt.Println("28 + 83 =", add(28, 83))
	fmt.Println("8 * 7 =", multiply(8, 7))
	fmt.Println("Hello World after swap: ", a, b)
}
