package main

import "fmt"

// Go has various value types including strings, integersm floats, booleans, etc. Here are a few basic examples:

func main() {
	// Strings can be added together with +
	fmt.Println("go" + "lang")
	// Integers and Floats
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// Booleans, with boolean operators as you'd expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
