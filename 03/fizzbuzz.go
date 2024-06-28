package main

import (
	"fmt" 
	"os"
	"strconv"
)

func main(){
	arg := os.Args[1]
	
	num, err := strconv.Atoi(arg)	

	if err != nil {
	fmt.Println("An error has occurred while trying to parse the first argument as an integer")
	return
	}

	if num % 2 == 0 {
		fmt.Println("fizz")
	} else {
	fmt.Println("buzz")
	}
}
