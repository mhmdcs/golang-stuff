package main

import "fmt"

func main() {
	a := 2
	b := 3.1
	
	fmt.Printf("a: %T %v\n", a, a)
	fmt.Printf("b: %8T %[1]v\n", b)

	a = int(b)
	fmt.Printf("a: %T %[1]v\n", a)
}
