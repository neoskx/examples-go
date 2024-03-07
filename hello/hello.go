package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("[hello] ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	// Get a greeting message and print it.
	// message, err := greetings.Hello("Gladys")
	messages, err := greetings.Hellos([]string{"Gladys", "Samantha", "Darrin"})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(messages)
}
