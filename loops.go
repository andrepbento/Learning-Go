package main

import (
    "fmt"
)

func main() {
    /*      for loop
        for [condition | (init; condition; increment) | Range]
        {
            statement(s);
        }
    */
    var b int = 15
    var a int

    numbers := [6]int{1, 2, 3, 5}

    for a := 0; a < 10; a++ {
        fmt.Printf("value of a: %d\n", a)
    }

    for a < b {
        a++
        fmt.Printf("value of a: %d\n", a)
    }

    for i, x := range numbers {
        fmt.Printf("value of x = %d at %d\n", x, i)
    }

    /*      nested for loop
    for [condition |  ( init; condition; increment ) | Range]
    {
        for [condition |  ( init; condition; increment ) | Range]
        {
            statement(s);
        }
        statement(s);
    }
    */
    var i, j int

    for i = 2; i < 100; i++ {
        for j = 2; j <= (i/j); j++ {
            if(i % j == 0) {
                break; // if factor found, not prime
            }
        }
        if(j > (i/j)) {
            fmt.Printf("%d is prime\n", i);
        }
    }

    /* Loop Control Statements
        break;
        continue;

        // goto statement
        goto label;
        ..
        .
        label: statement;
    */
    var c int = 10
    LOOP: for c < 20 {
        if c == 15 {
            c = c + 1
            goto LOOP
        }
        fmt.Printf("value of c: %d\n", c)
        c++
    }

    /*      The Infinite Loop
        for true {
            fmt.Printf("This loop will run forever.\n")
        }
    */
}