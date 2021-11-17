package main

import (
	"fmt"
	"myapp/doctor"
)

func main() {
	// var whatToSay string
	// whatToSay = "Hello Wold!"

	// whatToSay := "Hello world, again!" // Declaring the type given by the right side

	// sayHelloWorld(whatToSay)

	var whatToSay = doctor.Intro()

	fmt.Println(whatToSay)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
