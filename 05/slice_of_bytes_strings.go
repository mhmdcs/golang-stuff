package main

import "fmt"

func main() {
	s := []byte("string")

	fmt.Println(len(s), s)
	fmt.Println(s[3])
	fmt.Println(s[:2])
	fmt.Println(s[2:])
	fmt.Println(s[3:5], len(s[3:5]))
}
