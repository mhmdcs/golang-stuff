package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// run this program via `cat words.txt | go run find_popular_words.go` or `go run find_popular_words.go < words.txt`

func main() {
	scanner := bufio.NewScanner(os.Stdin) // create a scanner that scans standard input stream from the command line
	
	words := make(map[string]int) // make a non-nill map with empty storage, that uses strings as keys and ints as values
	
	// the default behavior for bufio's scanner is to scan input and use newlines \n as the delimiter
	// we will change this and configure bufio's scanner to use spaces as the delimiter, by passing bufio.ScanWords to Split()
	scanner.Split(bufio.ScanWords) // instead of giving us one line at a time, it will now give us one word at a time
	
	// remember, scanner.Scan() returns true, so this for-loop is basically a while true loop
	// scanner.Scan() will return false once ctrl+d is inputted to stdin stream, which signals end of input (EOF end of file)
	// scanner.Scan() will also return false if the program is fed a file via cat or piping, and the program reads (scans)
	// the file as an input stream, and then operating system itself will handle providing the EOF signal when the end of the file is reached
	for scanner.Scan() {
		// scanner.Text() is a string and we can immediately use that and index our map with it as the key
		// every time we see a word, we're gonna insert it in the map and if it's not there it'll be 1, and if it is there we increment further
		words[scanner.Text()]++
	}

	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var kvSlice []kv

	for k, v := range words {
		kvSlice = append(kvSlice, kv{k, v})
	}
	// now we've taken our words map and "unrolled" it into a slice that contains structs of keys and values
	// this will help us sort the keys and values as we want, we'll sort the values by the highest number
	
	// in Slice() second parameter, we pass in a callback function (called a literal function in go, and lambda/closure in others)
	// in the literal function, we are required to pass in two params i and j, and return a boolean
	// the i and j are index values in the slice, and then we return the boolean where we check that a kv's element val is greater and sort that way
	sort.Slice(kvSlice, func(i int, j int) bool {
		return kvSlice[i].val > kvSlice[j].val
	})

	// kvSlice[:3] means "show me the top three elements in the slice (element 0, 1, 2)
	for _, kv := range kvSlice[:3] {
		fmt.Println(kv.key, "appears", kv.val, "times")
	}

}

