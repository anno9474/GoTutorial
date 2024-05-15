package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")

	// print quote from other module
	fmt.Println(quote.Go())
}
