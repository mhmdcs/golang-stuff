package main

import ("fmt")

func main(){
	s := "élite"
	fmt.Printf("%T %[1]v %d\n", s, len(s)) 
	// the length of a string is the number of bytes needed to represent it in unicode utf-8 encoding, not the number of visual characters
	
	r := []rune(s)
	fmt.Printf("%T %[1]v %d\n", r, len(r))
	
	b := []byte(s)
	fmt.Printf("%T %[1]v %d\n", b, len(b))
}
