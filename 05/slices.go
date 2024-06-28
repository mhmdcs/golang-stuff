package main

import ("fmt")

func do(a [3]int, s []int) []int {
	// a = s // syntax error, can't assign a slice to an array
	
	a[0] = 4 // array will remain unchanged, arrays are passed by value
	s[0] = 4 // slice will change, slices are passed by reference
	
	c := make([]int, 5) // []init{0, 0, 0, 0}
	c[4] = 23
	copy(c, s) // copy the first 3 elements of the slice s into the new slice c
	
	return c
}

func main() {
	array := [...]int{1, 2, 3} // array of length 3
	slice := []int{9, 8, 7} // slice of length 3
	
	result := do (array, slice)
	fmt.Println(array, slice, result) // [1 2 3] [4 8 7] [4 8 7 0 23]
}
