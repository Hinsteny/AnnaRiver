package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Opt())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Hinsteny")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	// Request a greeting message.
	message1, err1 := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err1 != nil {
		log.Fatal(err1)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message1)
}
