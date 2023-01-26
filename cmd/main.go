package main

import (
	"flag"
	"fmt"
)

func main() {
        a := "aaaaaaaaaa"
        b := "aaaaaaaaaa"
        fmt.Println(a, b)
	// Basic command-line flag parsing
	exampleInt := flag.Int("n", 42, "help message for flag n")
	exampleStr := flag.String("example", "foo", "help message")
	flag.Parse()

	fmt.Println("n", *exampleInt)
	fmt.Println("example", *exampleStr)
}
