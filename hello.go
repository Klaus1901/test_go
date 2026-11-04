package main

import "fmt"

// Add function get two int params, and return their sum as int.
//
// Example link [https://www.mathsisfun.com/numbers/addition/html]
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Print(Add(2, 3))
}
