package main

import (
    "fmt"
)

// more info on.: https://blog.golang.org/defer-panic-and-recover

func main() {
    defer printTwo()
    printOne()
}

func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }