package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	message, err := greetings.Hello("takumines")
	if err != nil {
		log.SetPrefix("greetings: ")
		log.SetFlags(0)

		log.Fatal(err)
	}

	fmt.Println(message)
}
