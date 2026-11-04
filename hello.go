package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Add function get two int params, and return their sum as int.
//
// Example link [https://www.mathsisfun.com/numbers/addition/html]

type Number interface {
	constraints.Integer | constraints.Float
}

func Add[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(2, 3))
	fmt.Println(Add(2.1, 3.3))
}
