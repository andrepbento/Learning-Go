package main

import "fmt"

/*      @andrebento
    my first program in Go
*/

func main() {
    fmt.Println("Hello, World")

    fmt.Println("My first separated print!")

    // my first variable declaration in Go
    var age = 21
    fmt.Println(age)

    // some tests
    var _booleanVar = false
    fmt.Println(_booleanVar)
    _booleanVar = !_booleanVar
    fmt.Println(_booleanVar)

    int _int_age = 21
    age = age + 3
    fmt.Println("age = " + age)
}
