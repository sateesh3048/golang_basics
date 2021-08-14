package main

import "fmt"

//Features :-
// We can send multiple arguments as input to a function. if two variables of same type then we can declare type once.
// We can return multiple values from a function
// We can store a function inside a variable.
// We can send function as argument to another function.
// We can return function from a function.
// we can declare anonymous/nameless functions.
// We can achieve closures concept by using functions.
func main() {

	//concept:
	// We can store function inside a variable.
	// Since we can declare variable inside a function.
	// We can declare these functions inside function.
	increment := func(val int) int {
		val = val + 1
		return val
	}

	fmt.Println("increment: ", increment(10))

	result := sum(10, 20)
	fmt.Println("sum is: ", result)

	// Returning multiple values from a function.
	a, b := swap(23, 45)
	fmt.Println("swap is: ", a, b)

	// Sending function as argument to function
	fmt.Println(mathOp(10, 20, "sum", doMathOperations))

	nextId := nextSequence()
	//Calling closuers.
	fmt.Println("next id:", nextId())
}

//concept:
// function can take multiple values as input.
// if two arguments of the same type then we can declare them like a,b int
// We should mention return type of the function
func sum(a, b int) int {
	return a + b
}

//concept:
// We can return multiple values from function
func swap(a, b int) (int, int) {
	return b, a
}

//concept:
// We can send function as argument to the function
func mathOp(a, b int, operator string, doOperation func(int, int, string) int) int {
	result := doOperation(a, b, operator)
	return result
}

//concept:
// We can return function from the function
func nextSequence() func() int {
	count := 0
	return func() int {
		count = count + 1
		return count
	}
}

func doMathOperations(a, b int, operation string) int {
	switch operation {
	case "sum":
		return a + b
	case "sub":
		return a - b
	case "mul":
		return a * b
	case "div":
		return a / b
	default:
		return -1000000
	}
}
