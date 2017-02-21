package main

import (
    "fmt"
)

func main() {
    // type_name(expression)

    var sum int = 17
    var count int = 5
    var mean float32

    mean = float32(sum)/float32(count)
    fmt.Printf("Value of mean: %.2f\n", mean)
}