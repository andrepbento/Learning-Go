package main

import (
    "fmt"
)

const MAX int = 3

func main() {
    // & - access variable reference
    var a int = 10
    fmt.Printf("Address of a variable: %x\n", &a)
    //      Pointers
    // var var_name *var-type
    //  a) define a pointer variable
    //  b) assign the address of a variable to a pointer
    //  c) finally access the value at the adress avaiable in the pointer variable
    var b int = 20  // variable declaration
    var ptr *int    // pointer variable declaration

    ptr = &b        // store address of a in pointer variable

    fmt.Printf("\nAddress of b variable: %x\n", &b) 

    // Address stored in pointer variable
    fmt.Printf("Address stored in ptr variable: %x\n", ptr)

    // Access the value using the pointer
    fmt.Printf("Value of *ptr variable: %d\n", *ptr)

    //      nil Pointers
    var nil_ptr *int
    fmt.Printf("The value of nil_ptr is: %x\n", nil_ptr)

    if(nil_ptr == nil) {
        fmt.Printf("nil_ptr is nil")
    } else {
        fmt.Printf("nil_ptr is not nil")
    }

    //      Array of pointers
    c := []int{10,100,200}
    var ptr_array [MAX]*int

    for i := 0; i < MAX; i++ {
        ptr_array[i] = &c[i] // assign the address of integer
    }

    for i, element := range ptr_array {
        fmt.Printf("Value of c[%d] = %d\n", i, *element)
    }

    //      Pointer to pointer
    // var ptr_example **int
    var d int
    var ptr_d *int
    var pptr_d **int

    d = 3000

    // take the address of var
    ptr_d = &d

    // take the address of ptr_d using address of operator &
    pptr_d = &ptr_d

    // take the value using pptr_d
    fmt.Printf("Value of d = %d\n", d)
    fmt.Printf("Value avaiable at *ptr_d = %d\n", *ptr_d)
    fmt.Printf("Value avaiable at **pptr_d = %d\n", **pptr_d)

    //      Pointers to functions
    var e int = 100
    var f int = 200

    fmt.Printf("Before swap, value of e : %d\n", e)
    fmt.Printf("Before swap, value of f : %d\n", f)

    swap(&e, &f)

    fmt.Printf("After swap, value of e : %d\n", e)
    fmt.Printf("After swap, value of f : %d\n", f)
}

func swap(x *int, y *int) {
    var temp int
    temp = *x
    *x = *y
    *y = temp
}