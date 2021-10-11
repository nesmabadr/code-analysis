package main

import "fmt"

func main() {
	fmt.Println("main")
	fmt.Println("hello world")

	foo()
}

func foo() {
	fmt.Println("foo")

	fmt.Println("doing something")
	bar()
}

func bar() {
	fmt.Println("bar")
	fmt.Println("doing another thing")
}
