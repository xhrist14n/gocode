package main

import (
    "fmt"
    "greetings"
    "log"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    //message := greetings.Hello("Nadie")
    // Request a greeting message.
    message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}

