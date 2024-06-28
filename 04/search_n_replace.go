package main

import (
        "bufio"
        "fmt"
        "strings"
        "os"
)

// run this program with `cat strings.txt | go run search_and_replace.go mhmd wafa`
// or `go run search_and_replace.go mhmd wafa < strings.txt`

func main(){
        if len(os.Args) < 3 {
                fmt.Fprintln(os.Stderr, "Not enough args")
                os.Exit(-1)
        }

        old, new := os.Args[1], os.Args[2]

        scanner := bufio.NewScanner(os.Stdin) // scanner that scans from standard input

        // scanner.Scan() is a function that does the scanning, if it succeeds it returns a boolean true
        // so this is like a `for true {}` which is essentially a `while true {}` loop

        for scanner.Scan() {
                s := strings.Split(scanner.Text(), old) // scanner.Text() returns a single line
                t := strings.Join(s, new)

                fmt.Println(t)
        }
}
