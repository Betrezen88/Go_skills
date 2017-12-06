package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

<<<<<<< HEAD
//if function arguments are the same type you can omit type from all but the last.
func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println("28 + 83 =", add(28, 83))
	fmt.Println("8 * 7 =", multiply(8, 7))
=======
func main() {
	fmt.Println("28 + 83 =", add(28, 83))
>>>>>>> 83597d708c1cb6b526bd31f6f03d60a5955df1f6
}
