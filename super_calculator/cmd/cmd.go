package cmd

import (
    "flag"
    "fmt"
)

func init() {
    flag.StringVar(&greeting, "greeting", "Hello", "greeting message")
}

var greeting string

func Greet() {
    fmt.Println(greeting, "World!")
}
