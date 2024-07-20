package main

import "fmt"


func main() {
	// uint8 and byte are the same exact type, byte is just an alias for uint8, it just provides better readability 

	// Using uint8 gives the reader the impression that you're doing arithmetic operations
	var a uint8 = 10
	var b uint8 = 20
	var c uint8 = a + b
	
	fmt.Println("uint8 arithmetic: ", c) // Output: uint8 arithmetic: 30

	
	// Using byte for binary data gives the reader the impression that we're dealing with raw binary, like encoding or streaming
	data := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F}  // ASCII codes for "Hello"
	fmt.Println("byte data", data) // Output: byte data: [72 101 108 108 111]

	// Converting byte to string
	str := string(data)
	fmt.Println("string from bytes:", str)  // Output: string from bytes: Hello
}
