package main

import (
	"fmt"
	"os"
	"golangpractice/hello"
)

func main() {
	fmt.Println(hello.Say(os.Args[1:]))
}
