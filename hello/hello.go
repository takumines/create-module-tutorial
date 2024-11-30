package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	names := []string{"takumines", "taro", "jiro"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.SetPrefix("greetings: ")
		log.SetFlags(0)

		log.Fatal(err)
	}

	fmt.Println(messages)
}
