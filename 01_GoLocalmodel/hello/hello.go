package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	massage := greetings.Hello("Санёчек")
	fmt.Println(massage)
}