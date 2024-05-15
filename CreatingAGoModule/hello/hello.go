package main

import (
	"fmt"
	"github.com/Anno9474/GoTutorial/CreatingAGoModule/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Anno", "Dev", "CRO"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
