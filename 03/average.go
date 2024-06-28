package main

import (
	"fmt"
	"os"
	)

func main(){
	var sum float64
	var total int
	
	for {
		var input float64
		_, err := fmt.Fscanln(os.Stdin, &input) // pass address of input so Fscanln can use this pointer to assign std input value to input variable
		
		if err != nil {
		break // err exists and is io.EOF, meaning end of input/file, this is a signal to end input reading process, break the loop
		}
		
		sum += input
		total++
	}

	// signal end of input (EOF) to os.Stdin via typing ctrl+d in the console after running the program
	// or you could just read numbers from a file instead of inputting them on standard input
	// run `go run average.go < nums.txt` to redirect/feed the output of nums.txt as input to average.go
	// the end of file marker in nums.txt is used as the signal to end of input (like pressing ctrl+d on keyboard)
	// you could also `cat nums.txt | go run average.go` to pipe the output of catting nums.txt as input to average.go

	if total == 0 {
		fmt.Fprintln(os.Stderr, "No values were inputted")
		os.Exit(-1)
	}

	avg := sum / float64(total)
	//fmt.Sprintf("Sum is %v\nTotal is %f\nAverage is %f\n", sum, total, avg)
	fmt.Println("Average is", avg)
}
