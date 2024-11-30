package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("takumines")
	fmt.Println(message)
}
