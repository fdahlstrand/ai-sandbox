package main

import (
	"flag"
	"fmt"
)

func main() {
	namePtr := flag.String("name", "World", "Name to greet")
	flag.Parse()

	fmt.Printf("Hello, %s!\n", *namePtr)
	fmt.Println("End of the line")
}
