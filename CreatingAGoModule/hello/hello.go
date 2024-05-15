package main

import (
	"fmt"
	"github.com/Anno9474/GoTutorial/CreatingAGoModule/greetings"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
