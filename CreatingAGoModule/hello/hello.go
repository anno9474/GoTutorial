package main

import (
	"fmt"
	"github.com/Anno9474/GoTutorial/CreatingAGoModule/greetings"
)

func main() {
	message := greetings.Hello("Anno")
	fmt.Println(message)
}
