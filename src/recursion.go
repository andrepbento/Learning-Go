package main

import (
    "fmt"
)

var recursionCounter int = 0

func recursion() {
    recursionCounter++
    fmt.Println(recursionCounter,"- This recursive print is called by recursion() function")
    recursion()
}

func factorial(i int)int {
    if(i <= 1) {
        return 1
    }
    return i * factorial(i - 1)
}

func fibonaci(i int) (ret int) {
    if i == 0 {
        return 0
    }

    if i == 1 {
        return 1
    }

    return fibonaci(i - 1) + fibonaci(i - 2)
}

func main() {
    // recursion()

    var i int = 3
    fmt.Printf("Factorial of %d is %d\n", i, factorial(i))

    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", fibonaci(i))
    }
}