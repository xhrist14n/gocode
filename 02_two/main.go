package main

import (
    "fmt"
    "greetings"
)

func main() {
    message := greetings.Hello("Nadie")
    fmt.Println(message)
}

