package main

import (
    "fmt"
    "bufio"
    "os"
)

/*      @andrebento
    This program is an example of variables
    definition and usage in Go programming 
    lang
*/

func main() {
    fmt.Println("VARIABLES EXAMPLE PROGRAM\n")

    var i, j, k int
    fmt.Println("int variables:\n", i, j, k)

    var c, ch byte
    fmt.Println("byte variables:\n", c, ch)

    var f, salary float32
    salary = 3200.07
    fmt.Println("float32 variables:\n", f, salary)

    var d = 3; var g = 100; // declaration of d and g. Here d and g are int
    fmt.Println("variables without type:\n", d, g)

    var x float64
    x = 20.0
    fmt.Println(x)
    fmt.Printf("x is of type %T\n", x)

    y := 42
    fmt.Println(y)
    fmt.Printf("y is of type %T\n", y)

    var a, b, e = 3, 4, "gopher"
    fmt.Println(a)
    fmt.Printf("a is of type %T\n", a)
    fmt.Println(b)
    fmt.Printf("b is of type %T\n", b)
    fmt.Println(e)
    fmt.Printf("e is of type %T\n", e)
    
    var z = 20.0

    fmt.Println(z)
    fmt.Printf("z is of type %T\n", z)

    fmt.Print("Press 'Enter' to continue...")
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}
