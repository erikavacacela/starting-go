package main

func main() {
	// one way - declare, then assign (two steps)
	var firstNumber int
	firstNumber = 2
	println(firstNumber)

	// another way, declare type and name and assign value
	var secondNumber = 5
	println(secondNumber)

	// one step variable: declare name, assign value, and let Go figure out the type
	subtraction := 7
	println(subtraction)
}
