package main

import (
    "fmt"
    "math"
)

func myFirstFunction(param string) string {
    return "Function call "+param
}

func max(num1, num2 int) int {
    var result int
    if(num1 > num2) {
        result = num1
    } else {
        result = num2
    }
    return result
}

//  Returning multiple values from Function
func swapStrings(x, y string) (string, string) {
    return y, x
}

func swapByValue(x, y int) int {
    var temp int
    temp = x
    x = y
    y = temp
    return temp
}

func swapByReference(x *int, y *int) {
    var temp int
    temp = *x // save the value at address x
    *x = *y // put y into x
    *y = temp // put temp into y
}

func getSequence() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

/* define a circle */
type Circle struct {
    x, y, radius float64
}

/* define a method for circle */
func(circle Circle) area() float64 {
    return math.Pi * circle.radius * circle.radius
}

func main() {
    fmt.Println("FUNCTIONS PROGRAM EXAMPLE")
    // len() - function that returns the lenght of the type
    fmt.Printf("\nlen() function %d\n", len("go"))

    /*      function definition
        func function_name([parameter list]) [return_types]
        {
            body of the function
        }
    */
    fmt.Printf("%s\n", myFirstFunction("andrebento"))

    fmt.Printf("max(10, 100): %d\n", max(10, 100))

    name1, name2 := swapStrings("André", "Bento")
    fmt.Printf("swapStrings(\"André\", \"Bento\"): %s %s\n", name1, name2)

    // Call by value
    fmt.Printf("\nCall by value example\n")

    var a int = 100
    var b int = 200

    fmt.Printf("Before swap, value of a: %d\n", a)
    fmt.Printf("Before swap, value of b: %d\n", b)
    
    swapByValue(a, b)

    fmt.Printf("After swap, value of a: %d\n", a)
    fmt.Printf("After swap, value of b: %d\n", b)

    // Call by reference
    fmt.Printf("\nCall by reference example\n")

    var c int = 100
    var d int = 200

    fmt.Printf("Before swap, value of c: %d\n", c)
    fmt.Printf("Before swap, value of d: %d\n", d)
    
    swapByReference(&c, &d)

    fmt.Printf("After swap, value of c: %d\n", c)
    fmt.Printf("After swap, value of d: %d\n", d)

    // declare a function variable
    getSquareRoot := func(x float64) float64 {
        return math.Sqrt(x)
    }
    fmt.Printf("\nVariable as function: %.0f\n", getSquareRoot(9))

    // function closure example
    nextNumber := getSequence()

    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())

    nextNumber1 := getSequence()

    fmt.Println(nextNumber1())
    fmt.Println(nextNumber1())

    /* Method
        func (variable_name variable_data_type) function_name() [return_type] {
            function body
        }
    */
    circle := Circle{x: 0, y: 0, radius: 5}
    fmt.Printf("Circle area: %f", circle.area())
}
