package main

import "fmt"

func double(x int) int {
	return x + x
}

// both inputs are integers don't need to declare the title twice
func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet function")
}

func main() {
	greet()
	dozen := double(6)
	fmt.Println("a dozen is", dozen)
	bakersDozen := add(dozen, 1)
	fmt.Println("a bakers dozen is", bakersDozen)
	anotherBackerDozen := add(double(6), 1)
	fmt.Println("another bakers dozen is", anotherBackerDozen)
}
