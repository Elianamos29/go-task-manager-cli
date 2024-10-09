package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Gopher", "Your name")
	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)
}