package main

import (
    "fmt"
)

func main() {
    // Local variables are variables defined inside of a function
    // Global variables are variables defined outside of a function

    /*  Local variables takes preference between 
        themselves and Global varibles
    */

    /*      Formal Parameters
        They are threated as local variables with-in 
        that function and they will take preference
        over the global variables
    */

    /*      Initializing Local and 
                            Global variables
        int     -  0
        float32 -  0
        pointer -  nil
    */
    var integer int
    var realNumber float32
    var ptr *int
    fmt.Println(integer)
    fmt.Println(realNumber)
    if ptr == nil {
        fmt.Println("ptr is nil")
    }
    fmt.Println(ptr)
    fmt.Println(&ptr)
} 