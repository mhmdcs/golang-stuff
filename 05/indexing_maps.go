package main

import "fmt"

func main() {
	m := map[string]int{} // non-nill, but empty
	
	// attempting to index/get/read a value using a key that doesn't exist returns to us 0, which is the default zero value for int value
	a := m["the"] // returns 0
	fmt.Println(a)

	// the get function for map returns us two value actually, the first one is the value, the second is the status, this is a common pattern in go
	b, ok := m["and"] // returns 0, false
	fmt.Println(b, ok)

	m["the"]++ // we now incremented the 0 int default value for the "the" key, and now this key actually refers to a 1 int value
	
	c, ok := m["the"] // returns 1, true
	fmt.Println(c, ok)

	// if the "the" key refers to a valid value and ok is equal to true, then print success
	if d, ok := m["the"]; ok {
	fmt.Println("successfully read key", d)
	}
}
