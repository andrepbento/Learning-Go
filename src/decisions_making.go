package main

import (
    "fmt"
)

/*  @andrebento
    This program is an example on how to use
    decision makeing structures in Golang
*/

func main() {
    fmt.Println("DECISION MAKING PROGRAM EXAMPLE\n")
    /*      if statement
        if(boolean_expression)
        {
            statement(s) will execute if the boolean expression is true
        }
    */
    var a int = 10
    if(a < 20) {
        fmt.Printf("a is less than 20\n")
    }
    fmt.Printf("value of a is: %d\n", a)

    /*      if..else statement
        if(boolean_expression)
        {
            executes if boolean_expression is true
        }
        else
        {
            executes if boolean_expression is false
        }
    */
    if(true) {
        fmt.Println("Passei em if(true)")
    } else {
        fmt.Println("Passei em else")
    }

    var b int = 100
    if(b == 10) {
        fmt.Printf("Value of a is 10\n")
    } else if(b == 20) {
        fmt.Printf("Value of a is 20\n")
    } else {
        fmt.Printf("None of the values is matching\n")
    }

    /*      switch statement
        switch(boolean-expression or integral type){
            case boolean-expression or integral type  :
                statement(s);      
            case boolean-expression or integral type  :
                statement(s); 
            default : #Optional
                statement(s);
        }
    */
    var name string = "andrebento"
    switch(name) {
        case "andrebento":
            fmt.Println()
        case "andreacruz":
            fmt.Println("variable name is andreacruz")
        default:
            fmt.Println("none of cases is matching variable name")
    }

    /*      select statement
        select {
            case communication_clause:
                statement(s);
            case communication_clause:
                statement(s);
            default: #Optional
                statement(s);
        }
    */
    var c1, c2, c3 chan int
    var i1, i2 int
    select {
        case i1 = <-c1:
            fmt.Printf("received ", i1, " from c1\n")
        case c2 <- i2:
            fmt.Printf("sent ", i2, " to c2\n")
        case i3, ok := (<-c3):
            if ok {
                fmt.Printf("received ", i3, " from c3\n")
            } else {
                fmt.Printf("c3 is closed\n")
            }
        default:
            fmt.Printf("no communication\n")
    }

    // Just some variable declaration example
    v1 := 10
    fmt.Printf("%d", v1)
}