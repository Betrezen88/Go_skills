package main

import "fmt"

// var statement declares a list of variables, the type is last
var c, python, java bool

func main() {
	var i int
	// var declaration can include initializers, one per variable.
	var text, number, logic = "Happy", 42, true
	// var a, b = 10 <- Error: missing initializer for b.
	fmt.Println(i, c, python, java)
	fmt.Println(text, number, logic)
}
