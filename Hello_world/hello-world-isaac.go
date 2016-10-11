package main

// Our first program will print the classic “hello world” message. Here’s the full source code.

import "fmt"

func main() {
	fmt.Println("Hello world!")
}

// **** means run this in terminal ****

// To run the program, put the code in hello-world.go and use go run.
// **** $ go run hello-world.go ****

// Sometimes we’ll want to build our programs into binaries. We can do this using go build.
// **** $ go build hello-world.go ****

// We can then execute the built binary directly.
// **** $ ./hello-world ****
